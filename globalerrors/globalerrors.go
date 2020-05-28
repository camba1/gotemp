package globalerrors

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type ValidationError struct {
	Source      string
	FailureDesc []string
}

func (v *ValidationError) Error() string {
	var failureDesc string
	for _, desc := range v.FailureDesc {
		failureDesc += desc
	}
	return fmt.Sprintf("validation error in %s:\n %s ", v.Source, failureDesc)
}

type SrvError string

const (
	srvNoStartTxt               SrvError = "Unable to start %s server. Error: %v \n"
	srvNoHandlerTxt             SrvError = "Unable to register service handler. Error: %v"
	dbNoConnectionTxt           SrvError = "Unable to connect to DB %s. Error: %v\n"
	dbNoConnectionStringTxt     SrvError = "Unable to find DB connection string. Please set environment variable %s \n"
	dtProtoTimeStampToTimeStamp SrvError = "Unable to convert proto timestamp %v to timestamp. Error: %v \n"
	dtTimeStampToProtoTimeStamp SrvError = "Unable to convert timestamp %v to proto timestamp. Error: %v \n"
	dtInvalidValidityDates      SrvError = "The valid thru date (%v) must take place after the valid from date (%v)\n"
	missingField                SrvError = "%s must not be empty\n"
)

func (ge *SrvError) SrvNoStart(serviceName string, err error) string {
	return fmt.Sprintf(string(srvNoStartTxt), serviceName, err)
}

func (ge *SrvError) DbNoConnection(dbName string, err error) string {
	return fmt.Sprintf(string(dbNoConnectionTxt), dbName, err)
}

func (ge *SrvError) DbNoConnectionString(envVarName string) string {
	return fmt.Sprintf(string(dbNoConnectionStringTxt), envVarName)
}

func (ge *SrvError) SrvNoHandler(err error) string {
	return fmt.Sprintf(string(srvNoHandlerTxt), err)
}

func (ge *SrvError) DtProtoTimeStampToTimeStamp(currTimeStamp *timestamp.Timestamp, err error) string {
	return fmt.Sprintf(string(dtProtoTimeStampToTimeStamp), currTimeStamp, err)
}

func (ge *SrvError) DtTimeStampToProtoTimeStamp(currentTime time.Time, err error) string {
	return fmt.Sprintf(string(dtTimeStampToProtoTimeStamp), currentTime, err)
}

func (ge *SrvError) MissingField(fieldName string) string {
	return fmt.Sprintf(string(missingField), fieldName)
}

func (ge *SrvError) DtInvalidValidityDates(validFrom, validThru time.Time) string {
	return fmt.Sprintf(string(dtInvalidValidityDates), validFrom, validThru)
}
