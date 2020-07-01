package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	"goTemp/promotion/proto"
	"goTemp/promotion/server/statements"
	"log"
	"time"
)

//glErr: Holds the service global errors that are shared cross services
var glErr globalerrors.SrvError

//Promotion: Main entry point for promotion related services
type Promotion struct{}

//promoErr: Holds service specific errors
var promoErr statements.PromoErr

//UpdatePromotion: Updates a promotion based on the is provided in the inPromotion. Returns updated promotion
func (p *Promotion) UpdatePromotion(ctx context.Context, inPromotion *proto.Promotion, resp *proto.Response) error {
	_ = ctx

	outPromotion := &proto.Promotion{}

	if errVal := p.BeforeUpdatePromotion(ctx, inPromotion, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	sqlStatement := statements.SqlUpdate.String()

	convertedDates, err := globalUtils.TimeStampPPBToTime(inPromotion.GetValidFrom(), inPromotion.GetValidThru())
	if err != nil {
		return err
	}
	validFrom, validThru := convertedDates[0], convertedDates[1]

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
		log.Printf(promoErr.UpdateError(err))
		return err
	}

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
	if err != nil {
		return err
	}
	outPromotion.ValidFrom, outPromotion.ValidThru = convertedTimes[0], convertedTimes[1]
	resp.Promotion = outPromotion

	failureDesc, err := p.getAfterAlerts(ctx, outPromotion, "AfterUpdatePromotion")
	if err != nil {
		return err
	}
	resp.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	//if errVal := p.AfterUpdatePromotion(ctx, outPromotion, &proto.AfterFuncErr{}); errVal != nil {
	//	return errVal
	//}

	return nil
}

//DeletePromotion: Delete a promotion based on the promotion id in the searchId.id field. Returns number of affected promotions which should be one always
func (p *Promotion) DeletePromotion(ctx context.Context, searchid *proto.SearchId, resp *proto.Response) error {
	_ = ctx

	outPromotion := proto.Promotion{}
	if err := p.GetPromotionById(ctx, searchid, &outPromotion); err != nil {
		return err
	}
	if errVal := p.BeforeDeletePromotion(ctx, &outPromotion, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	sqlStatement := statements.SqlDelete.String()
	commandTag, err := conn.Exec(context.Background(), sqlStatement, searchid.Id)
	if err != nil {
		log.Printf(promoErr.DeleteError(searchid.Id, err))
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf(promoErr.DeleteRowNotFoundError(searchid.Id))
	}

	resp.AffectedCount = commandTag.RowsAffected()

	failureDesc, err := p.getAfterAlerts(ctx, &outPromotion, "AfterDeletePromotion")
	if err != nil {
		return err
	}
	resp.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}

	//if errVal := p.AfterDeletePromotion(ctx, &outPromotion, &proto.AfterFuncErr{}); errVal != nil {
	//	return errVal
	//}

	return nil
}

//GetPromotions: Returns a promotion slice based on the search parameters provided
func (p *Promotion) GetPromotions(ctx context.Context, searchParms *proto.SearchParams, promotions *proto.Promotions) error {

	_ = ctx

	sql := statements.SqlSelectAll.String()
	sqlWhereClause, values, err2 := p.buildSearchWhereClause(searchParms)
	if err2 != nil {
		return err2
	}

	sqlStatement := fmt.Sprintf(sql, sqlWhereClause, statements.MaxRowsToFetch)

	rows, err := conn.Query(context.Background(), sqlStatement, values...)

	if err != nil {
		log.Printf(promoErr.SelectReadError(err))
		return err
	}

	var validFrom time.Time
	var validThru time.Time
	for rows.Next() {
		var promo proto.Promotion
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

		convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
		if err != nil {
			return err
		}
		promo.ValidFrom, promo.ValidThru = convertedTimes[0], convertedTimes[1]

		promotions.Promotion = append(promotions.Promotion, &promo)
	}

	return nil
}

//buildSearchWhereClause: Builds a sql string to be used as the where clause in a sql statement. It also returns an interface
//slice with the values to be used as replacements in the sql statement. Currently only handles equality constraints, except
//for the date lookup which is done  as a contains clause
func (p *Promotion) buildSearchWhereClause(searchParms *proto.SearchParams) (string, []interface{}, error) {
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
	if searchParms.GetCustomerId() != "" {
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
		convertedDates, err := globalUtils.TimeStampPPBToTime(searchParms.GetValidDate())
		if err != nil {
			return "", nil, err
		}
		validFrom := convertedDates[0]
		sqlWhereClause += fmt.Sprintf(" AND promotion.validfrom <= $%d AND promotion.validthru >= $%d", i, i)
		values = append(values, validFrom)
		i++
	}
	return sqlWhereClause, values, nil
}

//GetPromotionById: Get a promotion for the given promotion id provided in searchId.id
func (p *Promotion) GetPromotionById(ctx context.Context, searchId *proto.SearchId, outPromotion *proto.Promotion) error {
	_ = ctx

	var validFrom time.Time
	var validThru time.Time

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

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
	if err != nil {
		return err
	}

	//Get the customer name
	getLookups(ctx, outPromotion)

	outPromotion.ValidFrom, outPromotion.ValidThru = convertedTimes[0], convertedTimes[1]

	return nil
}

//CreatePromotion: Creates a promotion based on the promotion passed in the inPromotion argument
func (p *Promotion) CreatePromotion(ctx context.Context, inPromotion *proto.Promotion, resp *proto.Response) error {
	_ = ctx

	outPromotion := &proto.Promotion{}

	if errVal := p.BeforeCreatePromotion(ctx, inPromotion, &proto.ValidationErr{}); errVal != nil {
		return errVal
	}

	convertedDates, err := globalUtils.TimeStampPPBToTime(inPromotion.GetValidFrom(), inPromotion.GetValidThru())
	if err != nil {
		return err
	}
	validFrom, validThru := convertedDates[0], convertedDates[1]

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

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru)
	if err != nil {
		return err
	}
	outPromotion.ValidFrom, outPromotion.ValidThru = convertedTimes[0], convertedTimes[1]

	resp.Promotion = outPromotion
	failureDesc, err := p.getAfterAlerts(ctx, outPromotion, "AfterCreatePromotion")
	if err != nil {
		return err
	}
	resp.ValidationErr = &proto.ValidationErr{FailureDesc: failureDesc}
	//
	//if errVal := p.AfterCreatePromotion(ctx, outPromotion, &proto.AfterFuncErr{}); errVal != nil {
	//	return errVal
	//}

	return nil
}

//getAfterAlerts: Call the appropriate after create/update/delete function and return the alert validation errors
//These alerts  are logged, but do not cause the record processing to fail
func (p *Promotion) getAfterAlerts(ctx context.Context, promotion *proto.Promotion, operation string) ([]string, error) {
	afterFuncErr := &proto.AfterFuncErr{}
	var errVal error
	if operation == "AfterDeletePromotion" {
		errVal = p.AfterDeletePromotion(ctx, promotion, afterFuncErr)
	}
	if operation == "AfterCreatePromotion" {
		errVal = p.AfterCreatePromotion(ctx, promotion, afterFuncErr)
	}
	if operation == "AfterUpdatePromotion" {
		errVal = p.AfterUpdatePromotion(ctx, promotion, afterFuncErr)
	}
	if errVal != nil {
		return []string{}, errVal
	}

	if len(afterFuncErr.GetFailureDesc()) > 0 {
		log.Printf("Alerts: %v: ", afterFuncErr.GetFailureDesc())
		return afterFuncErr.GetFailureDesc(), nil
	}
	return []string{}, nil
}

//getLookups: Function gets any read only lookups from other services (or cache) and add them to the ReadOnlyLookup field.
//In this case, it is pulling the customer name from the customer service based on the customerid
func getLookups(ctx context.Context, promotion *proto.Promotion) {
	lookup := idLookup{}
	custName, err := lookup.getCustomerName(ctx, promotion.CustomerId)
	if err != nil {
		log.Printf(glErr.MissingField("Customer Name"))
	} else {
		lookup := proto.ReadOnlyLookups{CustomerName: custName}
		promotion.ReadOnlyLookup = &lookup
	}

}
