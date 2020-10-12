package globalUtils

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	//AuditTopic defines a topic to be used when publishing and subscribing to audit related message in the broker
	AuditTopic = "Audit"
	//AuditQueueInsert is the queue to be used when retrieve audit messages from the broker via a subscription to allow multiple instances processing messages in parallel
	AuditQueueInsert = "Audit.Insert"
)

//auditMsg defines the structure of the audit messages to be sent to the broker
type auditMsg struct {
	topic        string
	objectToSend []byte
	header       AuditMsgHeader
}

// Header returns the header from the Audit Message
func (a auditMsg) Header() AuditMsgHeader {
	return a.header
}

//ObjectToSend returns the object that will be sent to the audit system in the Audit Message
func (a auditMsg) ObjectToSend() []byte {
	return a.objectToSend
}

//Topic returns the topic that will be used to store the object in the audit system
func (a auditMsg) Topic() string {
	return a.topic
}

//AuditMsgHeader defines the structure of the header portion of the audit message to be sent to the broker
type AuditMsgHeader map[string]string

// AuditMsgHeaderStruct is the struct version of the AuditMsgHeader (map) to allow for easier handling
type AuditMsgHeaderStruct struct {
	ServiceName  string
	ActionFunc   string
	ActionType   string
	ObjectId     string
	PerformedBy  int64
	ActionTime   time.Time
	ObjectName   string
	RecordedTime time.Time
	Id           int64
}

//GetId returns the id from audit header
func (a *AuditMsgHeaderStruct) GetId() int64 {
	return a.Id
}

//GetRecordedTime returns the recorded time from audit header
func (a *AuditMsgHeaderStruct) GetRecordedTime() time.Time {
	return a.RecordedTime
}

//GetObjectName returns the object name from audit header
func (a *AuditMsgHeaderStruct) GetObjectName() string {
	return a.ObjectName
}

//GetActionTime returns the  action time from the audit header
func (a *AuditMsgHeaderStruct) GetActionTime() time.Time {
	return a.ActionTime
}

//GetPerformedBy returns the  performed by  from the  audit header
func (a *AuditMsgHeaderStruct) GetPerformedBy() int64 {
	return a.PerformedBy
}

//GetObjectId returns the object id  from audit header
func (a *AuditMsgHeaderStruct) GetObjectId() string {
	return a.ObjectId
}

//GetActionType returns the action type from audit header
func (a *AuditMsgHeaderStruct) GetActionType() string {
	return a.ActionType
}

//GetActionFunc returns the action function name from audit header
func (a *AuditMsgHeaderStruct) GetActionFunc() string {
	return a.ActionFunc
}

//GetServiceName returns the service name from audit header
func (a *AuditMsgHeaderStruct) GetServiceName() string {
	return a.ServiceName
}

//AuditMsgHeaderStructs is the struct version of the AuditMsgHeader (map) to allow for easier handling
type AuditMsgHeaderStructs struct {
	Header []AuditMsgHeaderStruct
}

//AuditSearchParams defines the parameters to search for audit records
type AuditSearchParams struct {
	ObjectName      string
	ObjectId        string
	ActionTimeStart time.Time
	ActionTimeEnd   time.Time
}

//AuditSearchId is the parameter used to search audit records by id
type AuditSearchId struct {
	Id int64
}

//NewAuditMsg validates and create a new Audit message that is ready to be sent out to the broker
func NewAuditMsg(serviceName, actionFunc, actionType string, performedBy int64, objectName string, objectId string, objectToSend []byte) (*auditMsg, error) {
	var missingFields string
	if serviceName == "" {
		missingFields += " serviceName,"
	}
	if actionFunc == "" {
		missingFields += " actionFunc,"
	}
	if actionType == "" {
		missingFields += " actionType,"
	}
	if performedBy == 0 {
		missingFields += " performedBy,"
	}
	if objectId == "" {
		missingFields += " objectId,"
	}
	if objectToSend == nil {
		missingFields += " objectToSend,"
	}
	if missingFields != "" {
		return nil, fmt.Errorf("all fields must be filled in audit messages. The following fields are empty: %s",
			missingFields[1:len(missingFields)-1])
	}
	aud := auditMsg{
		topic:        AuditTopic,
		objectToSend: objectToSend,
		header: AuditMsgHeader{
			"service":     serviceName,
			"actionFunc":  actionFunc,
			"actionType":  actionType,
			"objectId":    objectId,
			"performedBy": strconv.FormatInt(performedBy, 10),
			"actionTime":  time.Now().Format(time.RFC3339),
			"objectName":  objectName,
		},
	}
	//log.Printf(" objid : %v, objidstr: %s ",  objectId, strconv.FormatInt(performedBy, 10))
	return &aud, nil
}

//AuditMsgHeaderToStruct converts the AuditMsgHeader to its struct based counterpart AuditMsgHeaderStruct
func AuditMsgHeaderToStruct(header AuditMsgHeader) (*AuditMsgHeaderStruct, error) {
	if header == nil {
		return nil, fmt.Errorf("message header cannot be nil when trying to convert to struct")
	}
	//objectId, err := strconv.ParseInt(header["objectId"], 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	//performedBy, err := euidToId(header["performedBy"])
	performedBy, err := strconv.ParseInt(header["performedBy"], 10, 64)
	if err != nil {
		return nil, err
	}
	actionTime, err := time.Parse(time.RFC3339, header["actionTime"])
	if err != nil {
		log.Printf("Unable to Format date %v, Error: %v\n", header["actionTime"], err)
		return nil, err
	}
	headerStruct := &AuditMsgHeaderStruct{
		ServiceName: header["service"],
		ActionFunc:  header["actionFunc"],
		ActionType:  header["actionType"],
		ObjectId:    header["objectId"],
		PerformedBy: performedBy,
		ActionTime:  actionTime,
		ObjectName:  header["objectName"],
	}

	return headerStruct, nil
}

//AuditSend converts a user to a byte array, compose an audit message and send that message to the broker for
//forwarding to the audit service
func AuditSend(ctx context.Context, mb MyBroker, serviceName, actionFunc, actionType string, objectName string, objectId string, byteMessage []byte) string {

	var auth AuthUtils
	performedBy, err := auth.GetCurrentUserFromContext(ctx)
	if err != nil {
		return "unable to get user from metadata"
	}

	auditMsg, err := NewAuditMsg(serviceName, actionFunc, actionType, performedBy, objectName, objectId, byteMessage)
	if err != nil {
		return glErr.AudFailureSending(actionType, objectId, err)
	}

	if err = mb.SendMsg(auditMsg.ObjectToSend(), auditMsg.Header(), auditMsg.Topic()); err != nil {
		return glErr.AudFailureSending(actionType, objectId, err)
	}
	return ""

}
