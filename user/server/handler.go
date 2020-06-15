package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	pb "goTemp/user/proto"
	"goTemp/user/server/statements"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

//glErr: Holds the service global errors that are shared cross services
var glErr globalerrors.SrvError

//Promotion: Main entry point for promotion related services
type User struct{}

//userErr: Holds service specific errors
var userErr statements.UserErr

//GetUserById: Get User from DB based on a given ID
func (u *User) GetUserById(ctx context.Context, searchId *pb.SearchId, outUser *pb.User) error {
	_ = ctx

	var validFrom time.Time
	var validThru time.Time
	var createDate time.Time
	var updateDate time.Time

	sqlStatement := statements.SqlSelectById.String()
	err := conn.QueryRow(context.Background(), sqlStatement,
		searchId.GetId()).
		Scan(
			&outUser.Id,
			&outUser.Firstname,
			&outUser.Lastname,
			&validFrom,
			&validThru,
			&outUser.Active,
			&outUser.Pwd,
			&outUser.Name,
			&outUser.Email,
			&outUser.Company,
			&createDate,
			&updateDate,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		} else {
			log.Printf(userErr.SelectRowReadError(err))
			return err
		}

	}

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru, createDate, updateDate)
	if err != nil {
		return err
	}
	outUser.ValidFrom, outUser.ValidThru = convertedTimes[0], convertedTimes[1]
	outUser.Createdate, outUser.Updatedate = convertedTimes[2], convertedTimes[3]

	return nil
}

//GetUsers: Search the Users table in the DB based in a set of search parameters
func (u *User) GetUsers(ctx context.Context, searchParms *pb.SearchParams, users *pb.Users) error {

	_ = ctx

	values, sqlStatement, err2 := u.getSQLForSearch(searchParms)
	if err2 != nil {
		return err2
	}

	rows, err := conn.Query(context.Background(), sqlStatement, values...)

	if err != nil {
		log.Printf(userErr.SelectReadError(err))
		return err
	}

	var validFrom time.Time
	var validThru time.Time
	var createDate time.Time
	var updateDate time.Time

	for rows.Next() {
		var user pb.User
		err := rows.
			Scan(
				&user.Id,
				&user.Firstname,
				&user.Lastname,
				&validFrom,
				&validThru,
				&user.Active,
				&user.Pwd,
				&user.Name,
				&user.Email,
				&user.Company,
				&createDate,
				&updateDate,
			)
		if err != nil {
			log.Printf(userErr.SelectScanError(err))
			return err
		}

		convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru, createDate, updateDate)
		if err != nil {
			return err
		}
		user.ValidFrom, user.ValidThru = convertedTimes[0], convertedTimes[1]
		user.Createdate, user.Updatedate = convertedTimes[2], convertedTimes[3]

		users.User = append(users.User, &user)
	}

	return nil
}

//getSQLForSearch: Combine the where clause built in the buildSearchWhereClause method with the rest of the sql
//statement to return the final search for users sql statement
func (u *User) getSQLForSearch(searchParms *pb.SearchParams) ([]interface{}, string, error) {
	sql := statements.SqlSelectAll.String()
	sqlWhereClause, values, err := u.buildSearchWhereClause(searchParms)
	if err != nil {
		return nil, "", err
	}

	sqlStatement := fmt.Sprintf(sql, sqlWhereClause, statements.MaxRowsToFetch)
	return values, sqlStatement, nil
}

//buildSearchWhereClause: Builds a sql string to be used as the where clause in a sql statement. It also returns an interface
//slice with the values to be used as replacements in the sql statement. Currently only handles equality constraints, except
//for the date lookup which is done  as a contains clause
func (u *User) buildSearchWhereClause(searchParms *pb.SearchParams) (string, []interface{}, error) {
	sqlWhereClause := " where 1=1"
	var values []interface{}

	i := 1
	if searchParms.GetId() != 0 {
		sqlWhereClause += fmt.Sprintf(" AND appuser.id = $%d", i)
		values = append(values, searchParms.GetId())
		i++
	}
	if searchParms.GetFisrtname() != "" {
		sqlWhereClause += fmt.Sprintf(" AND appuser.firstname = $%d", i)
		values = append(values, searchParms.GetFisrtname())
		i++
	}
	if searchParms.GetLastname() != "" {
		sqlWhereClause += fmt.Sprintf(" AND appuser.lastname = $%d", i)
		values = append(values, searchParms.GetLastname())
		i++
	}
	if searchParms.GetEmail() != "" {
		sqlWhereClause += fmt.Sprintf(" AND appuser.email = $%d", i)
		values = append(values, searchParms.GetEmail())
		i++
	}
	if searchParms.GetCompany() != "" {
		sqlWhereClause += fmt.Sprintf(" AND appuser.company = $%d", i)
		values = append(values, searchParms.GetCompany())
		i++
	}
	if searchParms.GetValidDate() != nil {
		convertedDates, err := globalUtils.TimeStampPPBToTime(searchParms.GetValidDate())
		if err != nil {
			return "", nil, err
		}
		validFrom := convertedDates[0]
		sqlWhereClause += fmt.Sprintf(" AND appuser.validfrom <= $%d AND appuser.validthru >= $%d", i, i)
		values = append(values, validFrom)
		i++
	}
	return sqlWhereClause, values, nil
}

//CreateUser: Creates a user in the Database, including hashing their password before saving it. Calls before and after create functions
//for validations and sending record to the audit service via the broker
func (u *User) CreateUser(ctx context.Context, inUser *pb.User, resp *pb.Response) error {
	_ = ctx
	outUser := &pb.User{}

	if errVal := u.BeforeCreateUser(ctx, inUser, &pb.ValidationErr{}); errVal != nil {
		return errVal
	}

	hashedPwd, err := u.hashpwd(inUser.GetPwd())
	if err != nil {
		return err
	}
	inUser.Pwd = hashedPwd

	convertedDates, err := globalUtils.TimeStampPPBToTime(inUser.GetValidFrom(), inUser.GetValidThru())
	if err != nil {
		return err
	}
	validFrom, validThru := convertedDates[0], convertedDates[1]

	var createDate, updateDate time.Time

	sqlStatement := statements.SqlInsert.String()
	errIns := conn.QueryRow(context.Background(), sqlStatement,
		inUser.GetFirstname(),
		inUser.GetLastname(),
		validFrom,
		validThru,
		inUser.GetActive(),
		inUser.GetPwd(),
		inUser.GetEmail(),
		inUser.GetCompany(),
	).
		Scan(
			&outUser.Id,
			&outUser.Firstname,
			&outUser.Lastname,
			&validFrom,
			&validThru,
			&outUser.Active,
			&outUser.Pwd,
			&outUser.Name,
			&outUser.Email,
			&outUser.Company,
			&createDate,
			&updateDate,
		)

	if errIns != nil {
		log.Printf(userErr.InsertError(err))
		return errIns
	}

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru, createDate, updateDate)
	if err != nil {
		return err
	}
	outUser.ValidFrom, outUser.ValidThru = convertedTimes[0], convertedTimes[1]
	outUser.Createdate, outUser.Updatedate = convertedTimes[2], convertedTimes[3]

	resp.User = outUser
	failureDesc, err := u.getAfterAlerts(ctx, outUser, "AfterCreateUser")
	if err != nil {
		return err
	}
	resp.ValidationErr = &pb.ValidationErr{FailureDesc: failureDesc}

	return nil
}

//UpdateUser: Update a user in the Database. Calls before and after create functions
//for validations and sending record to the audit service via the broker
func (u *User) UpdateUser(ctx context.Context, inUser *pb.User, resp *pb.Response) error {
	_ = ctx

	outUser := &pb.User{}
	if errVal := u.BeforeUpdateUser(ctx, inUser, &pb.ValidationErr{}); errVal != nil {
		return errVal
	}

	sqlStatement := statements.SqlUpdate.String()

	convertedDates, err := globalUtils.TimeStampPPBToTime(inUser.GetValidFrom(), inUser.GetValidThru())
	if err != nil {
		return err
	}
	validFrom, validThru := convertedDates[0], convertedDates[1]

	var createDate time.Time
	updateDate := time.Now()
	fmt.Printf("updatedate: %v\n", updateDate)

	err = conn.QueryRow(context.Background(), sqlStatement,
		inUser.GetFirstname(),
		inUser.GetLastname(),
		validFrom,
		validThru,
		inUser.GetActive(),
		inUser.GetPwd(),
		inUser.GetEmail(),
		inUser.GetCompany(),
		updateDate,
		inUser.GetId(),
	).Scan(
		&outUser.Id,
		&outUser.Firstname,
		&outUser.Lastname,
		&validFrom,
		&validThru,
		&outUser.Active,
		&outUser.Pwd,
		&outUser.Name,
		&outUser.Email,
		&outUser.Company,
		&createDate,
		&updateDate,
	)
	if err != nil {
		log.Printf(userErr.UpdateError(err))
		return err
	}

	convertedTimes, err := globalUtils.TimeToTimeStampPPB(validFrom, validThru, createDate, updateDate)
	if err != nil {
		return err
	}
	outUser.ValidFrom, outUser.ValidThru = convertedTimes[0], convertedTimes[1]
	outUser.Createdate, outUser.Updatedate = convertedTimes[2], convertedTimes[3]

	resp.User = outUser
	failureDesc, err := u.getAfterAlerts(ctx, outUser, "AfterUpdateUser")
	if err != nil {
		return err
	}
	resp.ValidationErr = &pb.ValidationErr{FailureDesc: failureDesc}

	return nil
}

//DeleteUser: Delete a user in the Database based on the user ID. Calls before and after create functions
//for validations and sending record to the audit service via the broker
func (u *User) DeleteUser(ctx context.Context, searchid *pb.SearchId, resp *pb.Response) error {
	_ = ctx

	outUser := pb.User{}
	if err := u.GetUserById(ctx, searchid, &outUser); err != nil {
		return err
	}
	if errVal := u.BeforeDeleteUser(ctx, &outUser, &pb.ValidationErr{}); errVal != nil {
		return errVal
	}

	sqlStatement := statements.SqlDelete.String()
	commandTag, err := conn.Exec(context.Background(), sqlStatement, searchid.Id)
	if err != nil {
		log.Printf(userErr.DeleteError(searchid.Id, err))
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf(userErr.DeleteRowNotFoundError(searchid.Id))
	}

	resp.AffectedCount = commandTag.RowsAffected()

	failureDesc, err := u.getAfterAlerts(ctx, &outUser, "AfterDeleteUser")
	if err != nil {
		return err
	}
	resp.ValidationErr = &pb.ValidationErr{FailureDesc: failureDesc}

	return nil
}

//getAfterAlerts: Call the appropriate after create/update/delete function and return the alert validation errors
//These alerts  are logged, but do not cause the record processing to fail
func (u *User) getAfterAlerts(ctx context.Context, user *pb.User, operation string) ([]string, error) {
	afterFuncErr := &pb.AfterFuncErr{}
	var errVal error
	if operation == "AfterDeleteUser" {
		errVal = u.AfterDeleteUser(ctx, user, afterFuncErr)
	}
	if operation == "AfterCreateUser" {
		errVal = u.AfterCreateUser(ctx, user, afterFuncErr)
	}
	if operation == "AfterUpdateUser" {
		errVal = u.AfterUpdateUser(ctx, user, afterFuncErr)
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

//Auth: Authenticate user and return a new JWT token
func (u *User) Auth(ctx context.Context, user *pb.User, token *pb.Token) error {
	_ = ctx

	searchParams := pb.SearchParams{
		Email: user.Email,
	}
	outUsers := pb.Users{}
	if err := u.GetUsers(ctx, &searchParams, &outUsers); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(outUsers.User[0].GetPwd()), []byte(user.GetPwd())); err != nil {
		return err
	}

	ts := TokenService{}
	encodeUser := &pb.User{
		Id:      outUsers.User[0].Id,
		Active:  outUsers.User[0].Active,
		Company: outUsers.User[0].Company,
	}
	tokenString, err := ts.Encode(encodeUser)
	if err != nil {
		return err
	}

	token.Token = tokenString
	token.Valid = false

	return nil
}

//GetUsersByEmail: Get a user given an email address. Internally just calls GetUsers.
func (u *User) GetUsersByEmail(ctx context.Context, searchString *pb.SearchString, outUsers *pb.Users) error {
	searchParams := pb.SearchParams{
		Email: searchString.Value,
	}
	err := u.GetUsers(ctx, &searchParams, outUsers)
	if err != nil {
		return err
	}
	return nil
}

//hashpwd: Hash a plan text string
func (u *User) hashpwd(plainPwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwd), err
}

//ValidateToken: Validate token to ensure user is authenticated
func (u *User) ValidateToken(ctx context.Context, inToken *pb.Token, outToken *pb.Token) error {
	_ = ctx
	ts := TokenService{}
	claims, err := ts.Decode(inToken.Token)
	if err != nil {
		return err
	}
	if claims == nil {
		return fmt.Errorf(glErr.AuthNilClaim(serviceName))
	}
	if claims.User.Id == 0 || claims.Issuer != ClaimIssuer {
		//fmt.Printf("claim User %v", claims.User)
		return fmt.Errorf(glErr.AuthInvalidClaim(serviceName))
	}
	//fmt.Printf("Claim User %v", claims.User)
	outToken.Token = inToken.Token
	outToken.Valid = true
	return nil

}

//userIdFromToken: Return the user id from the token
func (u *User) userIdFromToken(ctx context.Context, inToken *pb.Token) (int64, error) {
	_ = ctx
	if inToken.Valid == false {
		return 0, fmt.Errorf(glErr.AuthInvalidClaim(serviceName))
	}
	ts := TokenService{}
	claims, err := ts.Decode(inToken.Token)
	if err != nil {
		return 0, err
	}
	if claims.User.Id == 0 {
		//fmt.Printf("claim User %v", claims.User)
		return 0, fmt.Errorf(glErr.AuthInvalidClaim(serviceName))
	}
	return claims.User.Id, nil

}
