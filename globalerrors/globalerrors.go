package globalerrors

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

//ValidationError: Custom error type Used to compile multiple validations errors as one error
type ValidationError struct {
	Source      string
	FailureDesc []string
}

//Error: Combine all the message stored in the ValidationError.FailureDesc slice and return it as one
func (v *ValidationError) Error() string {
	var failureDesc string
	for _, desc := range v.FailureDesc {
		failureDesc += desc
	}
	return fmt.Sprintf("validation error in %s:\n %s ", v.Source, failureDesc)
}

//SrvError: String type relating to all the global errors
type SrvError string

const (
	srvNoStartTxt               SrvError = "Unable to start %s server. Error: %v \n"
	srvNoHandlerTxt             SrvError = "Unable to register service handler. Error: %v"
	dbNoConnectionTxt           SrvError = "Unable to connect to DB %s. Error: %v\n"
	dbNoConnectionStringTxt     SrvError = "Unable to find DB connection string. Please set environment variable %s \n"
	dbConnectRetry              SrvError = "Attempting to connect to DB again. Retry number: %d. Previous error: %v\n"
	dtProtoTimeStampToTimeStamp SrvError = "Unable to convert proto timestamp %v to timestamp. Error: %v \n"
	dtTimeStampToProtoTimeStamp SrvError = "Unable to convert timestamp %v to proto timestamp. Error: %v \n"
	dtInvalidValidityDates      SrvError = "The valid thru date (%v) must take place after the valid from date (%v)\n"
	missingField                SrvError = "%s must not be empty\n"
	authNoMetaData              SrvError = "Unable to read meta-date for end point: %s\n"
	authInvalidToken            SrvError = "invalid token\n"
	authNilToken                SrvError = "Invalid nil user token\n"
	authNilClaim                SrvError = "invalid nil %s claim\n"
	authInvalidClaim            SrvError = "Invalid %s claim\n"
	authNoUserInToken           SrvError = "unable to get logged in user from metadata. Error: %v\n"
	brkBadMarshall              SrvError = "Unable to marshall object %v. Error: %v\n"
	brkNoMessageSent            SrvError = "Unable to send message to broker for topic %s. Error: %v\n"
	brkNoConnection             SrvError = "Unable to connect to broker: Error: %v\n"
	brkUnableToSetSubs          SrvError = "Unable to setup broker subscription for topic %s. Error: %v\n"
	brkBadUnMarshall            SrvError = "Unable to unmarshall received object for topic %s. Message received: %v. Error: %v\n"
	audFailureSending           SrvError = "Unable to send %s audit information for record %s. Error: %v\n"
)

const (
	marshalFullMap          SrvError = "Unable to marshal full map. Error: %v\n"
	unMarshalByteFullMap    SrvError = "Unable to unmarshal byte full version map to struct. Error: %v\n"
	marshalPartialMap       SrvError = "Unable to marshal partial  map. Error: %v\n"
	unMarshalBytePartialMap SrvError = "Unable to unmarshal byte partial version of map to proto struct. Error: %v\n"
)

const (
	cacheUnableToWrite       SrvError = "Unable to write to cache with key %s. Error: %v\n"
	cacheDBNameNotSet        SrvError = "Cache Database Name is not set. Please provide a value\n"
	cacheUnableToReadVal     SrvError = "Unable to read key %v. Error: %v\n"
	cacheUnableToDeleteVal   SrvError = "Unable to delete key %v from cache. Error %v\n"
	cacheTooManyValuesToList SrvError = "Requested too many keys to list from cache. Max number is %d\n"
	cacheListError           SrvError = "Unable to list cache Keys .Error %v\n"
)

/*
Functions that return the error message formatted with the information passed in as arguments to the individual functions
*/

func (ge *SrvError) SrvNoStart(serviceName string, err error) string {
	return fmt.Sprintf(string(srvNoStartTxt), serviceName, err)
}

func (ge *SrvError) DbNoConnection(dbName string, err error) string {
	return fmt.Sprintf(string(dbNoConnectionTxt), dbName, err)
}

func (ge *SrvError) DbNoConnectionString(envVarName string) string {
	return fmt.Sprintf(string(dbNoConnectionStringTxt), envVarName)
}

func (ge *SrvError) DbConnectRetry(RetryNum int, err error) string {
	return fmt.Sprintf(string(dbConnectRetry), RetryNum, err)
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

func (ge *SrvError) AuthNoMetaData(endpoint string) string {
	return fmt.Sprintf(string(authNoMetaData), endpoint)
}

func (ge *SrvError) AuthInvalidToken() string {
	return fmt.Sprintf(string(authInvalidToken))
}

func (ge *SrvError) AuthNilToken() string {
	return fmt.Sprintf(string(authNilToken))
}

func (ge *SrvError) AuthNilClaim(claimType string) string {
	return fmt.Sprintf(string(authNilClaim), claimType)
}

func (ge *SrvError) AuthInvalidClaim(claimType string) string {
	return fmt.Sprintf(string(authInvalidClaim), claimType)
}

func (ge *SrvError) BrkBadMarshall(objToMarshal string, err error) string {
	return fmt.Sprintf(string(brkBadMarshall), objToMarshal, err)
}

func (ge *SrvError) BrkNoMessageSent(objToMarshal string, err error) string {
	return fmt.Sprintf(string(brkNoMessageSent), objToMarshal, err)
}

func (ge *SrvError) BrkNoConnection(err error) string {
	return fmt.Sprintf(string(brkNoConnection), err)
}

func (ge *SrvError) BrkUnableToSetSubs(topic string, err error) string {
	return fmt.Sprintf(string(brkUnableToSetSubs), topic, err)
}

func (ge *SrvError) BrkBadUnMarshall(topic string, message []byte, err error) string {
	return fmt.Sprintf(string(brkBadUnMarshall), topic, message, err)
}

func (ge *SrvError) AudFailureSending(operation string, id string, err error) string {
	return fmt.Sprintf(string(audFailureSending), operation, id, err)
}

func (ge *SrvError) AuthNoUserInToken(err error) string {
	return fmt.Sprintf(string(authNoUserInToken), err)
}

func (ge *SrvError) MarshalFullMap(err error) string {
	return fmt.Sprintf(string(marshalFullMap), err)
}
func (ge *SrvError) UnMarshalByteFullMap(err error) string {
	return fmt.Sprintf(string(unMarshalByteFullMap), err)
}
func (ge *SrvError) MarshalPartialMap(err error) string {
	return fmt.Sprintf(string(marshalPartialMap), err)
}
func (ge *SrvError) UnMarshalBytePartialMap(err error) string {
	return fmt.Sprintf(string(unMarshalBytePartialMap), err)
}

func (ge *SrvError) CacheUnableToWrite(key string, err error) string {
	return fmt.Sprintf(string(cacheUnableToWrite), key, err)
}

func (ge *SrvError) CacheDBNameNotSet() string {
	return fmt.Sprintf(string(cacheDBNameNotSet))
}

func (ge *SrvError) CacheUnableToReadVal(key string, err error) string {
	return fmt.Sprintf(string(cacheUnableToReadVal), key, err)
}

func (ge *SrvError) CacheUnableToDeleteVal(key string, err error) string {
	return fmt.Sprintf(string(cacheUnableToDeleteVal), key, err)
}

//
func (ge *SrvError) CacheTooManyValuesToList(maxValues int) string {
	return fmt.Sprintf(string(cacheTooManyValuesToList), maxValues)
}

//cacheListError
func (ge *SrvError) CacheListError(err error) string {
	return fmt.Sprintf(string(cacheListError), err)
}
