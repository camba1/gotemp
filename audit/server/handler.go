package main

import (
	"context"
	"github.com/micro/go-micro/v2/broker"
	"goTemp/audit/server/statements"
	"goTemp/globalUtils"
	"log"
	"time"
)

//mb: Broker instance to send/receive message from pub/sub system
var mb globalUtils.MyBroker

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

	//actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, objectDetail
	outHeaderStruct := globalUtils.AuditMsgHeaderStruct{}

	log.Printf("Storing Message in DB: %v", headerStruct.ActionType)

	var b1 time.Time

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
		//"\\\\",
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
			&b1,
		)

	if errIns != nil {
		//log.Printf(userErr.InsertError(err))
		return errIns
	}
	//log.Printf("Saved audit record: %v\n", outHeaderStruct)
	log.Printf("Saved audit record \n")
	return nil
}
