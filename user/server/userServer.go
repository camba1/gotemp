package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"goTemp/globalUtils"
	pb "goTemp/user/proto"
	"log"
	"os"
	"strings"
)

// serviceName service identifier
const serviceName = "goTemp.api.user"

// const serviceName = "user"

const (
	// dbName Name of the DB hosting the data
	dbName = "appuser"
	// dbConStrEnvVarName Name of Environment variable that contains connection string to DB
	dbConStrEnvVarName = "POSTGRES_CONNECT"
)

// Other constants
const (
	DisableAuditRecordsEnvVarName = "DISABLE_AUDIT_RECORDS"
)

//  conn Database connection
var conn *pgx.Conn

// glDisableAuditRecords Allows all insert,update,delete records to be sent out to the broker for forwarding to
var glDisableAuditRecords = false

// AuthWrapper defines the authentication middleware
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		// User login is excepted from authentication
		if req.Endpoint() == "UserSrv.Auth" || req.Endpoint() == "UserSrv.CreateUser" {
			return fn(ctx, req, resp)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return fmt.Errorf(glErr.AuthNoMetaData(req.Endpoint()))
		}

		auth, ok := meta["Authorization"]
		if !ok {
			return fmt.Errorf(glErr.AuthNilToken())
		}
		authSplit := strings.SplitAfter(auth, " ")
		if len(authSplit) != 2 {
			return fmt.Errorf(glErr.AuthNilToken())
		}
		token := authSplit[1]

		// token := meta["Token"]

		log.Printf("endpoint: %v", req.Endpoint())

		// Validate token
		var u User
		outToken := &pb.Token{}
		err := u.ValidateToken(ctx, &pb.Token{Token: token}, outToken)
		if err != nil {
			return err
		}
		if outToken.Valid != true {
			return fmt.Errorf(glErr.AuthInvalidToken())
		}

		// Add current user to context to use in saving audit records
		// userId, err := u.userIdFromToken(ctx, outToken)
		// if err != nil {
		// 	return fmt.Errorf("unable to get user id from token for endpoint %v\n", req.Endpoint())
		// }
		// ctx2 := metadata.Set(ctx, "userid", strconv.FormatInt(userId, 10))

		if outToken.EUid == "" {
			return fmt.Errorf("unable to get user id from token for endpoint %v\n", req.Endpoint())
		}
		ctx2 := metadata.Set(ctx, "userid", outToken.EUid)

		return fn(ctx2, req, resp)
	}
}

// getDBConnString gets the connection string to the DB
func getDBConnString() string {
	connString := os.Getenv(dbConStrEnvVarName)
	if connString == "" {
		log.Fatalf(glErr.DbNoConnectionString(dbConStrEnvVarName))
	}
	return connString
}

// connectToDB calls the Util pgxDBConnect to connect to the database. Service will try to connect a few times
// before giving up and throwing an error
func connectToDB() *pgx.Conn {
	var pgxConnect globalUtils.PgxDBConnect
	dbConn, err := pgxConnect.ConnectToDBWithRetry(dbName, getDBConnString())
	if err != nil {
		log.Fatalf(glErr.DbNoConnection(dbName, err))
	}
	return dbConn
}

// loadConfig loads service configuration
func loadConfig() {
	// conf, err := config.NewConfig()
	// if err != nil {
	// 	log.Fatalf("Unable to create new application configuration object. Err: %v\n", err)
	// 	// log.Fatal(err)
	// }
	// defer conf.Close()
	//
	// src := env.NewSource()
	//
	// err = conf.Load(src)
	// // ws, err := src.Read()
	// if err != nil {
	// 	log.Fatalf("Unable to load application configuration object. Err: %v\n", err)
	// 	// log.Fatal(err)
	// }
	// test := conf.Map()
	// // log.Printf("conf %v\n", ws.Data)
	//
	// log.Printf("conf map %v\n", test)

	audits := os.Getenv(DisableAuditRecordsEnvVarName)
	if audits == "true" {
		glDisableAuditRecords = true
	} else {
		glDisableAuditRecords = false
	}
}

func main() {

	// TODO: get version number from external source
	// setup metrics collector
	metricsWrapper := newMetricsWrapper()

	// instantiate service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.WrapHandler(AuthWrapper),
		micro.WrapHandler(metricsWrapper),
	)

	service.Init()
	err := pb.RegisterUserSrvHandler(service.Server(), new(User))
	if err != nil {
		log.Fatalf(glErr.SrvNoHandler(err))
	}

	// Load configuration
	loadConfig()

	// Connect to DB
	conn = connectToDB()

	defer conn.Close(context.Background())

	// setup the nats broker

	mb.Br = service.Options().Broker
	defer mb.Br.Disconnect()

	// Initialize http server for metrics export
	go runHttp()

	//  Run Service
	err = service.Run()
	if err != nil {
		log.Fatalf(glErr.SrvNoStart(serviceName, err))
	}

}
