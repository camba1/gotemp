package globalUtils

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	AuditTopic       = "Audit"
	AuditQueueInsert = "Audit.Insert"
)

type auditMsg struct {
	topic        string
	objectToSend []byte
	header       AuditMsgHeader
}

func (a auditMsg) Header() AuditMsgHeader {
	return a.header
}

func (a auditMsg) ObjectToSend() []byte {
	return a.objectToSend
}

func (a auditMsg) Topic() string {
	return a.topic
}

type AuditMsgHeader map[string]string

type AuditMsgHeaderStruct struct {
	ServiceName string
	ActionFunc  string
	ActionType  string
	ObjectId    int64
	PerformedBy int64
	ActionTime  time.Time
	ObjectName  string
}

func (a *AuditMsgHeaderStruct) GetObjectName() string {
	return a.ObjectName
}

func (a *AuditMsgHeaderStruct) GetActionTime() time.Time {
	return a.ActionTime
}

func (a *AuditMsgHeaderStruct) GetPerformedBy() int64 {
	return a.PerformedBy
}

func (a *AuditMsgHeaderStruct) GetObjectId() int64 {
	return a.ObjectId
}

func (a *AuditMsgHeaderStruct) GetActionType() string {
	return a.ActionType
}

func (a *AuditMsgHeaderStruct) GetActionFunc() string {
	return a.ActionFunc
}

func (a *AuditMsgHeaderStruct) GetServiceName() string {
	return a.ServiceName
}

func NewAuditMsg(serviceName, actionFunc, actionType string, performedBy int64, objectName string, objectId int64, objectToSend []byte) (*auditMsg, error) {
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
	if objectId == 0 {
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
			"objectId":    strconv.FormatInt(objectId, 10),
			"performedBy": strconv.FormatInt(performedBy, 10),
			"actionTime":  time.Now().Format(DateLayoutISO),
			"objectName":  objectName,
		},
	}
	//log.Printf(" objid : %v, objidstr: %s ",  objectId, strconv.FormatInt(performedBy, 10))
	return &aud, nil
}

func AuditMsgHeaderToStruct(header AuditMsgHeader) (*AuditMsgHeaderStruct, error) {
	if header == nil {
		return nil, fmt.Errorf("message header cannot be nil when trying to convert to struct")
	}
	objectId, err := strconv.ParseInt(header["objectId"], 10, 64)
	if err != nil {
		return nil, err
	}
	performedBy, err := strconv.ParseInt(header["performedBy"], 10, 64)
	if err != nil {
		return nil, err
	}
	actionTime, err := time.Parse(DateLayoutISO, header["actionTime"])
	if err != nil {
		log.Printf("Unable to Format date %v\n", header["actionTime"])
	}
	headerStruct := &AuditMsgHeaderStruct{
		ServiceName: header["service"],
		ActionFunc:  header["actionFunc"],
		ActionType:  header["actionType"],
		ObjectId:    objectId,
		PerformedBy: performedBy,
		ActionTime:  actionTime,
		ObjectName:  header["objectName"],
	}

	return headerStruct, nil
}
