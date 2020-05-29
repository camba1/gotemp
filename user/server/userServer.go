package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/micro/go-micro/v2"
	pb "goTemp/user/proto"
	"log"
	"os"
)

//serviceName: service identifier
const serviceName = "user"

const (
	//dbName: Name of the DB hosting the data
	dbName = "appuser"
	//dbConStrEnvVarName: Name of Environment variable that contains connection string to DB
	dbConStrEnvVarName = "POSTGRES_CONNECT"
)

// conn: Database connection
var conn *pgx.Conn

//getDBConnString: Get the connection string to the DB
func getDBConnString() string {
	connString := os.Getenv(dbConStrEnvVarName)
	if connString == "" {
		log.Fatalf(glErr.DbNoConnectionString(dbConStrEnvVarName))
	}
	return connString
}

func main() {

	//instantiate service
	service := micro.NewService(
		micro.Name(serviceName),
	)
	service.Init()
	err := pb.RegisterUserSrvHandler(service.Server(), new(User))
	if err != nil {
		log.Fatalf(glErr.SrvNoHandler(err))
	}

	// Connect to DB
	conn, err = pgx.Connect(context.Background(), getDBConnString())
	if err != nil {
		log.Fatalf(glErr.DbNoConnection(dbName, err))
	}
	defer conn.Close(context.Background())

	// Run Service
	if err := service.Run(); err != nil {
		log.Fatalf(glErr.SrvNoStart(serviceName, err))
	}

}
