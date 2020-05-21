package globalerrors

import "fmt"

type SrvError string

const (
	srvNoStartTxt           SrvError = "Unable to start %s server. Error: %v \n"
	srvNoHandlerTxt         SrvError = "Unable to register service handler. Error: %v"
	DbNoConnectionTxt       SrvError = "Unable to connect to DB %s. Error: %v\n"
	DbNoConnectionStringTxt SrvError = "Unable to find DB connection string. Please set environment variable %s \n"
)

func (ge *SrvError) SrvNoStart(serviceName string, err error) string {
	return fmt.Sprintf(string(srvNoStartTxt), serviceName, err)
}

func (ge *SrvError) DbNoConnection(dbName string, err error) string {
	return fmt.Sprintf(string(DbNoConnectionTxt), dbName, err)
}

func (ge *SrvError) DbNoConnectionString(envVarName string) string {
	return fmt.Sprintf(string(DbNoConnectionStringTxt), envVarName)
}

func (ge *SrvError) SrvNoHandler(err error) string {
	return fmt.Sprintf(string(srvNoHandlerTxt), err)
}
