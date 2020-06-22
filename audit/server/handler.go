package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/micro/go-micro/v2/broker"
	"goTemp/audit/server/statements"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	"log"
)

//mb: Broker instance to send/receive message from pub/sub system
var mb globalUtils.MyBroker

//auditErr: Holds service specific errors
var auditErr statements.UserErr

//AuditSrv: Struct that will hold the audit service functionality
type AuditSrv struct{}

//SubsToBrokerMsg: Subscribes to the pub/sub broker to receive messages
func (a *AuditSrv) SubsToBrokerInsertMsg() error {
	err := mb.SubToMsg(a.processBrokerMessage, globalUtils.AuditTopic, globalUtils.AuditQueueInsert)
	if err != nil {
		return err
	}
	return nil
}

//processBrokerMessage: extract the topic, header and message payload from the message received from the broker
func (a *AuditSrv) processBrokerMessage(p broker.Event) error {
	log.Printf("Extracting received message")
	topic, header, message, err := mb.GetMsg(p)
	if err != nil {
		return err
	}
	log.Printf("Converting received message")
	headerStruct, err := globalUtils.AuditMsgHeaderToStruct(header)
	if err != nil {
		return err
	}
	log.Printf("headerStruct : %v\n topic: %s\n", headerStruct, topic)

	err = a.createAudit(topic, headerStruct, message)
	if err != nil {
		return err
	}
	return nil
}

//createAudit: Insert audit information into the database
func (a *AuditSrv) createAudit(topic string, headerStruct *globalUtils.AuditMsgHeaderStruct, message []byte) error {

	outHeaderStruct := globalUtils.AuditMsgHeaderStruct{}

	log.Printf("Storing Message in DB: %v", headerStruct.ActionType)

	//var b1 time.Time

	sqlStatement := statements.SqlInsert.String()

	errIns := conn.QueryRow(context.Background(), sqlStatement,
		headerStruct.GetActionTime(),
		topic,
		headerStruct.GetServiceName(),
		headerStruct.GetActionFunc(),
		headerStruct.GetActionType(),
		headerStruct.GetObjectName(),
		headerStruct.GetObjectId(),
		headerStruct.GetPerformedBy(),
		message,
	).
		Scan(
			&outHeaderStruct.ActionTime,
			&topic,
			&outHeaderStruct.ServiceName,
			&outHeaderStruct.ActionFunc,
			&outHeaderStruct.ActionType,
			&outHeaderStruct.ObjectName,
			&outHeaderStruct.ObjectId,
			&outHeaderStruct.PerformedBy,
			&outHeaderStruct.RecordedTime,
		)

	if errIns != nil {
		//log.Printf(userErr.InsertError(err))
		return errIns
	}
	//log.Printf("Saved audit record: %v\n", outHeaderStruct)
	log.Printf("Saved audit record \n")
	return nil
}

//GetAudit: Get both header and message for an audit record
func (a *AuditSrv) GetAudit(ctx context.Context, searchId *globalUtils.AuditSearchId) (*globalUtils.AuditMsgHeaderStruct, *[]byte, error) {
	_ = ctx
	headerStruct := globalUtils.AuditMsgHeaderStruct{}
	message := make([]byte, 0)
	topic := ""

	sqlStatement := statements.SqlSelectById.String()
	err := conn.QueryRow(context.Background(), sqlStatement,
		searchId.Id).
		Scan(
			&headerStruct.ActionTime,
			&topic,
			&headerStruct.ServiceName,
			&headerStruct.ActionFunc,
			&headerStruct.ActionType,
			&headerStruct.ObjectName,
			&headerStruct.ObjectId,
			&headerStruct.PerformedBy,
			&headerStruct.RecordedTime,
			&headerStruct.Id,
			&message,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &headerStruct, &message, nil
		} else {
			log.Printf(auditErr.SelectRowReadError(err))
			return nil, nil, err
		}

	}
	return &headerStruct, &message, nil
}

//GetAudits: Retrieve audit headers from the DB. The details are not pulled as they could grow the message size significantly
func (a *AuditSrv) GetAudits(ctx context.Context, searchParms *globalUtils.AuditSearchParams) (*globalUtils.AuditMsgHeaderStructs, error) {
	_ = ctx
	var headerStructs globalUtils.AuditMsgHeaderStructs

	if _, err := a.validateSearchParams(searchParms); err != nil {
		return nil, err
	}

	values, sqlStatement, err2 := a.getSQLForSearch(searchParms)
	if err2 != nil {
		return nil, err2
	}

	rows, err := conn.Query(context.Background(), sqlStatement, values...)

	if err != nil {
		log.Printf(auditErr.SelectReadError(err))
		return nil, err
	}

	for rows.Next() {
		var headerStruct globalUtils.AuditMsgHeaderStruct
		var topic string
		err := rows.
			Scan(
				&headerStruct.ActionTime,
				&topic,
				&headerStruct.ServiceName,
				&headerStruct.ActionFunc,
				&headerStruct.ActionType,
				&headerStruct.ObjectName,
				&headerStruct.ObjectId,
				&headerStruct.PerformedBy,
				&headerStruct.RecordedTime,
				&headerStruct.Id,
			)
		if err != nil {
			log.Printf(auditErr.SelectScanError(err))
			return nil, err
		}
		headerStructs.Header = append(headerStructs.Header, headerStruct)
	}
	return &headerStructs, err

}

//validateSearchParams: Check if we have a valid search params argument. All params are mandatory to avoid
//potentially retrieve a large number of records
func (a *AuditSrv) validateSearchParams(searchParms *globalUtils.AuditSearchParams) ([]string, error) {
	var FailureDesc []string
	if searchParms.ObjectName == "" {
		FailureDesc = append(FailureDesc, glErr.MissingField("Object name"))
	}
	if searchParms.ObjectId == "" {
		FailureDesc = append(FailureDesc, glErr.MissingField("Object id"))
	}
	if searchParms.ActionTimeStart.IsZero() {
		FailureDesc = append(FailureDesc, glErr.MissingField("Action time Start"))
	}
	if searchParms.ActionTimeEnd.IsZero() {
		FailureDesc = append(FailureDesc, glErr.MissingField("Action time End"))
	}
	if searchParms.ActionTimeStart.After(searchParms.ActionTimeEnd) || searchParms.ActionTimeStart.Equal(searchParms.ActionTimeEnd) {
		FailureDesc = append(FailureDesc, glErr.DtInvalidValidityDates(searchParms.ActionTimeStart, searchParms.ActionTimeEnd))
	}
	if len(FailureDesc) > 0 {
		return FailureDesc, &globalerrors.ValidationError{Source: "validateSearchParams", FailureDesc: FailureDesc}
	}
	return FailureDesc, nil
}

//getSQLForSearch: Combine the where clause built in the buildSearchWhereClause method with the rest of the sql
//statement to return the final search for users sql statement
func (a *AuditSrv) getSQLForSearch(searchParms *globalUtils.AuditSearchParams) ([]interface{}, string, error) {
	sql := statements.SqlSelectAll.String()
	sqlWhereClause, values, err := a.buildSearchWhereClause(searchParms)
	if err != nil {
		return nil, "", err
	}

	sqlStatement := fmt.Sprintf(sql, sqlWhereClause, statements.MaxRowsToFetch)
	return values, sqlStatement, nil
}

//buildSearchWhereClause: Builds a sql string to be used as the where clause in a sql statement. It also returns an interface
//slice with the values to be used as replacements in the sql statement. Currently only handles equality constraints, except
//for the date lookup which is done  as a contains clause
func (a *AuditSrv) buildSearchWhereClause(searchParms *globalUtils.AuditSearchParams) (string, []interface{}, error) {
	sqlWhereClause := " where 1=1"
	var values []interface{}

	i := 1
	if searchParms.ObjectName != "" {
		sqlWhereClause += fmt.Sprintf(" AND audit.objectname = $%d", i)
		values = append(values, searchParms.ObjectName)
		i++
	}
	if searchParms.ObjectId != "" {
		sqlWhereClause += fmt.Sprintf(" AND audit.objectid = $%d", i)
		values = append(values, searchParms.ObjectId)
		i++
	}
	if !searchParms.ActionTimeStart.IsZero() {
		sqlWhereClause += fmt.Sprintf(" AND audit.actiontime >= $%d", i)
		values = append(values, searchParms.ActionTimeStart)
		i++
	}
	if !searchParms.ActionTimeEnd.IsZero() {
		sqlWhereClause += fmt.Sprintf(" AND audit.actiontime <= $%d", i)
		values = append(values, searchParms.ActionTimeEnd)
		i++
	}
	return sqlWhereClause, values, nil
}
