package globalUtils

import "fmt"

const topic = "Audit"

type auditMsg struct {
	topic        string
	objectToSend []byte
	header       auditMsgHeader
}

func (a auditMsg) Header() auditMsgHeader {
	return a.header
}

func (a auditMsg) ObjectToSend() []byte {
	return a.objectToSend
}

func (a auditMsg) Topic() string {
	return a.topic
}

type auditMsgHeader map[string]string

func NewAuditMsg(serviceName, actionFunc, actionType string, performedBy, objectId int64, objectToSend []byte) (*auditMsg, error) {
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
		topic:        topic,
		objectToSend: objectToSend,
		header: auditMsgHeader{
			"service":     serviceName,
			"actionFunc":  actionFunc,
			"actionType":  actionType,
			"objectId":    string(objectId),
			"performedBy": string(performedBy),
		},
	}
	return &aud, nil
}
