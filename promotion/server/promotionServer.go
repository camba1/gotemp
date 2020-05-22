package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/jackc/pgx/v4"
	"github.com/micro/go-micro/v2"
	"goTemp/globalerrors"
	pb "goTemp/promotion"
	"goTemp/promotion/server/statements"
	"log"
	"os"
	"time"
)

//serviceName: service identifier
const serviceName = "promotion"

const (
	//dbName: Name of the DB hosting the data
	dbName = "postgres"
	//dbConStrEnvVarName: Name of Environment variable that contains connection string to DB
	dbConStrEnvVarName = "POSTGRES_CONNECT"
)

// conn: Database connection
var conn *pgx.Conn

//glErr: Holds the service global errors that are shared cross services
var glErr globalerrors.SrvError

//promoErr: Holds service specific errors
var promoErr statements.PromoErr

type Promotion struct{}

//UpdatePromotion: Updates a promotion based on the is provided in the inPromotion. Returns updated promotion
func (p *Promotion) UpdatePromotion(ctx context.Context, inPromotion *pb.Promotion, outPromotion *pb.Promotion) error {
	_ = ctx
	//sqlStatement := ` update promotion set
	//					 name = $1,
	//					 description = $2,
	//					 validfrom = $3,
	//					 validthru = $4,
	//					 customerid = $5,
	//					 active = $6,
	//					 approvalstatus = $7,
	//					 prevapprovalstatus = $8
	//					where id = $9
	//					RETURNING id, name, description, validfrom, validthru, customerid, active, approvalstatus,  prevapprovalstatus
	//				`

	sqlStatement := statements.SqlUpdate.String()

	validFrom, err := ptypes.Timestamp(inPromotion.GetValidFrom())
	if err != nil {
		log.Printf("Unable to convert valid from proto timestamp to timestamp when trying to update promotion. Error: %v \n", err)
		return err
	}

	validThru, err := ptypes.Timestamp(inPromotion.GetValidThru())
	if err != nil {
		log.Printf("Unable to convert valid thru proto timestamp to timestamp  when trying to update promotion. Error: %v \n", err)
		return err
	}

	err = conn.QueryRow(context.Background(), sqlStatement,
		inPromotion.GetName(),
		inPromotion.GetDescription(),
		validFrom,
		validThru,
		inPromotion.GetCustomerId(),
		inPromotion.GetActive(),
		inPromotion.GetApprovalStatus(),
		inPromotion.GetPrevApprovalStatus(),
		inPromotion.GetId(),
	).Scan(
		&outPromotion.Id,
		&outPromotion.Name,
		&outPromotion.Description,
		&validFrom,
		&validThru,
		&outPromotion.CustomerId,
		&outPromotion.Active,
		&outPromotion.ApprovalStatus,
		&outPromotion.PrevApprovalStatus,
	)
	if err != nil {
		//log.Printf("Unable to update promotion. Error: %v \n", err)
		log.Printf(promoErr.UpdateError(err))
		return err
	}
	outPromotion.ValidFrom, _ = ptypes.TimestampProto(validFrom)
	outPromotion.ValidThru, _ = ptypes.TimestampProto(validThru)

	return nil
}

//DeletePromotion: Delete a promotion based on the promotion id in the searchId.id field. Returns number of affected promotions which should be one always
func (p *Promotion) DeletePromotion(ctx context.Context, searchid *pb.SearchId, affectedCount *pb.AffectedCount) error {
	_ = ctx
	sqlStatement := statements.SqlDelete.String() //"delete from promotion where id = $1"
	commandTag, err := conn.Exec(context.Background(), sqlStatement, searchid.Id)
	if err != nil {
		//log.Printf("Unable to delete promotion %v. Error: %v\n", searchid.Id, err)
		log.Printf(promoErr.DeleteError(searchid.Id, err))
		return err
	}
	if commandTag.RowsAffected() != 1 {
		//return fmt.Errorf("row with id %d not found. Unable to delete the row", searchid.Id)
		return fmt.Errorf(promoErr.DeleteRowNotFoundError(searchid.Id))
	}

	affectedCount.Value = commandTag.RowsAffected()
	return nil
}

//GetPromotions: Returns a promotion slice based on the search parameters provided
func (p *Promotion) GetPromotions(ctx context.Context, searchParms *pb.SearchParams, promotions *pb.Promotions) error {

	_ = ctx

	//sqlStatement := `select
	//					id, name, description, validfrom, validthru,
	//					customerid, active, approvalstatus,  prevapprovalstatus
	//				 from promotion`

	sqlStatement := statements.SqlSelectAll.String()
	sqlWhereClause, values, err2 := p.buildSearchWhereClause(searchParms)
	if err2 != nil {
		return err2
	}

	sqlStatement += sqlWhereClause

	rows, err := conn.Query(context.Background(), sqlStatement, values...)

	if err != nil {
		log.Printf(promoErr.SelectReadError(err))
		return err
	}

	var validFrom time.Time
	var validThru time.Time
	for rows.Next() {
		var promo pb.Promotion
		err := rows.
			Scan(
				&promo.Id,
				&promo.Name,
				&promo.Description,
				&validFrom,
				&validThru,
				&promo.CustomerId,
				&promo.Active,
				&promo.ApprovalStatus,
				&promo.PrevApprovalStatus,
			)
		if err != nil {
			log.Printf(promoErr.SelectScanError(err))
			return err
		}
		promo.ValidFrom, _ = ptypes.TimestampProto(validFrom)
		promo.ValidThru, _ = ptypes.TimestampProto(validThru)
		promotions.Promotion = append(promotions.Promotion, &promo)
	}

	return nil
}

//buildSearchWhereClause: Builds a sql string to be used as the where clause in a sql statement. It also returns an interface
//slice with the values to be used as replacements in the sql statement. Currently only handles equality constraints, except
//for the date lookup which is done  as a contains clause
func (p *Promotion) buildSearchWhereClause(searchParms *pb.SearchParams) (string, []interface{}, error) {
	sqlWhereClause := " where 1=1"
	var values []interface{}

	i := 1
	if searchParms.GetId() != 0 {
		sqlWhereClause += fmt.Sprintf(" AND promotion.id = $%d", i)
		values = append(values, searchParms.Id)
		i++
	}
	if searchParms.GetName() != "" {
		sqlWhereClause += fmt.Sprintf(" AND promotion.name = $%d", i)
		values = append(values, searchParms.Name)
		i++
	}
	if searchParms.GetCustomerId() != 0 {
		sqlWhereClause += fmt.Sprintf(" AND promotion.customerid = $%d", i)
		values = append(values, searchParms.CustomerId)
		i++
	}
	if searchParms.GetProductId() != 0 {
		sqlWhereClause += fmt.Sprintf(" AND promotion.productid = $%d", i)
		values = append(values, searchParms.ProductId)
		i++
	}
	if searchParms.GetValidDate() != nil {
		validFrom, err := ptypes.Timestamp(searchParms.GetValidDate())
		if err != nil {
			log.Printf("Unable to convert valid from proto timestamp to timestamp when trying to search promotions. Error: %v \n", err)
			return "", nil, err
		}
		sqlWhereClause += fmt.Sprintf(" AND promotion.validfrom <= $%d AND promotion.validthru >= $%d", i, i)
		values = append(values, validFrom)
		i++
	}
	return sqlWhereClause, values, nil
}

//GetPromotionById: Get a promotion for the given promotion id provided in searchId.id
func (p *Promotion) GetPromotionById(ctx context.Context, searchId *pb.SearchId, outPromotion *pb.Promotion) error {
	_ = ctx

	var validFrom time.Time
	var validThru time.Time

	//sqlStatement := `select
	//					id, name, description, validfrom, validthru,
	//					customerid, active, approvalstatus,  prevapprovalstatus
	//				from promotion
	//				where id = $1`
	sqlStatement := statements.SqlSelectById.String()
	err := conn.QueryRow(context.Background(), sqlStatement,
		searchId.Id).
		Scan(
			&outPromotion.Id,
			&outPromotion.Name,
			&outPromotion.Description,
			&validFrom,
			&validThru,
			&outPromotion.CustomerId,
			&outPromotion.Active,
			&outPromotion.ApprovalStatus,
			&outPromotion.PrevApprovalStatus,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		} else {
			log.Printf(promoErr.SelectRowReadError(err))
			return err
		}

	}

	outPromotion.ValidFrom, _ = ptypes.TimestampProto(validFrom)
	outPromotion.ValidThru, _ = ptypes.TimestampProto(validThru)

	return nil
}

//CreatePromotion: Creates a promotion based on the promotion passed in the inPromotion argument
func (p *Promotion) CreatePromotion(ctx context.Context, inPromotion *pb.Promotion, outPromotion *pb.Promotion) error {
	_ = ctx

	validFrom, err := ptypes.Timestamp(inPromotion.GetValidFrom())
	if err != nil {
		log.Printf("Unable to convert valid from proto timestamp to timestamp when trying to create promotion. Error: %v \n", err)
		return err
	}

	validThru, err := ptypes.Timestamp(inPromotion.GetValidThru())
	if err != nil {
		log.Printf("Unable to convert valid thru proto timestamp to timestamp  when trying to create promotion. Error: %v \n", err)
		return err
	}

	//sqlStatement := `insert into promotion (name, description, validfrom, validthru, customerid, active, approvalstatus, prevapprovalstatus)
	//			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	//			RETURNING id, name, description, validfrom, validthru, customerid, active, approvalstatus,  prevapprovalstatus `

	sqlStatement := statements.SqlInsert.String()
	errIns := conn.QueryRow(context.Background(), sqlStatement,
		inPromotion.GetName(),
		inPromotion.GetDescription(),
		validFrom,
		validThru,
		inPromotion.GetCustomerId(),
		inPromotion.GetActive(),
		inPromotion.GetApprovalStatus(),
		inPromotion.GetPrevApprovalStatus(),
	).
		Scan(
			&outPromotion.Id,
			&outPromotion.Name,
			&outPromotion.Description,
			&validFrom,
			&validThru,
			&outPromotion.CustomerId,
			&outPromotion.Active,
			&outPromotion.ApprovalStatus,
			&outPromotion.PrevApprovalStatus,
		)

	if errIns != nil {
		log.Printf(promoErr.InsertError(err))
		return errIns
	}

	outPromotion.ValidFrom, _ = ptypes.TimestampProto(validFrom)
	outPromotion.ValidThru, _ = ptypes.TimestampProto(validThru)

	return nil
}

//getDBConnString: Get the connection string to the DB
func getDBConnString() string {
	connString := os.Getenv(dbConStrEnvVarName)
	if connString == "" {
		log.Fatalf(glErr.DbNoConnectionString(dbConStrEnvVarName))
	}
	return connString
}

//func testwhere() {
//	var p Promotion
//
//	priceVTtime, _ := time.Parse("2006-01-02", "2021-05-24")
//	validDate, _ := ptypes.TimestampProto(priceVTtime)
//	search := pb.SearchParams{
//		Id:         123,
//		Name:       "123",
//		ProductId:  123,
//		CustomerId: 123,
//		ValidDate:  validDate,
//	}
//	a,b,err := p.buildSearchWhereClause(&search)
//	if err != nil {
//		fmt.Printf("error: %v\n", err)
//	}
//	fmt.Printf("sql: %s\n", a)
//	fmt.Printf("values: %v\n", b)
//
//}

func main() {

	//testwhere()

	//instantiate service
	service := micro.NewService(
		micro.Name(serviceName),
	)
	service.Init()
	err := pb.RegisterPromotionSrvHandler(service.Server(), new(Promotion))
	if err != nil {
		log.Fatalf(glErr.SrvNoHandler(err))
	}

	// Connect to DB
	conn, err = pgx.Connect(context.Background(), getDBConnString())
	if err != nil {
		log.Fatalf(glErr.DbNoConnection(dbName, err))
	}
	defer conn.Close(context.Background())

	// Run Service
	if err := service.Run(); err != nil {
		log.Fatalf(glErr.SrvNoStart(serviceName, err))
	}

}
