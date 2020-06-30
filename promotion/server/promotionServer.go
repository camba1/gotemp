package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	store2 "github.com/micro/go-micro/v2/store"
	"goTemp/globalUtils"
	"goTemp/promotion/proto"
	pb "goTemp/user/proto"
	userSrv "goTemp/user/proto"
	"log"
	"os"
	"time"
)

//serviceName: service identifier
const serviceName = "promotion"

const (
	//dbName: Name of the DB hosting the data
	dbName = "postgres"
	//dbConStrEnvVarName: Name of Environment variable that contains connection string to DB
	dbConStrEnvVarName = "POSTGRES_CONNECT"
)

//Other constants
const (
	DisableAuditRecordsEnvVarName = "DISABLE_AUDIT_RECORDS"
)

// conn: Database connection
var conn *pgx.Conn

//enableAuditRecords: Allows all insert,update,delete records to be sent out to the broker for forwarding to
var glDisableAuditRecords = false

//AuthWrapper: Authentication middleware
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return fmt.Errorf(glErr.AuthNoMetaData(req.Endpoint()))
		}

		token := meta["Token"]
		log.Printf("endpoint: %v", req.Endpoint())

		userClient := userSrv.NewUserSrvService("user", client.DefaultClient)
		outToken, err := userClient.ValidateToken(ctx, &pb.Token{Token: token})
		if err != nil {
			return err
		}
		if outToken.Valid != true {
			return fmt.Errorf(glErr.AuthInvalidToken())
		}

		if outToken.EUid == "" {
			return fmt.Errorf("unable to get user id from token for endpoint %v\n", req.Endpoint())
		}
		ctx2 := metadata.Set(ctx, "userid", outToken.EUid)

		return fn(ctx2, req, resp)
	}
}

//getDBConnString: Get the connection string to the DB
func getDBConnString() string {
	connString := os.Getenv(dbConStrEnvVarName)
	if connString == "" {
		log.Fatalf(glErr.DbNoConnectionString(dbConStrEnvVarName))
	}
	return connString
}

//connectToDB: Call the Util pgxDBConnect to connect to the database. Service will try to connect a few times
//before giving up and throwing an error
func connectToDB() *pgx.Conn {
	var pgxConnect globalUtils.PgxDBConnect
	dbConn, err := pgxConnect.ConnectToDBWithRetry(dbName, getDBConnString())
	if err != nil {
		log.Fatalf(glErr.DbNoConnection(dbName, err))
	}
	return dbConn
}

func loadConfig() {
	//conf, err := config.NewConfig()
	//if err != nil {
	//	log.Fatalf("Unable to create new application configuration object. Err: %v\n", err)
	//	//log.Fatal(err)
	//}
	//defer conf.Close()
	//
	//src := env.NewSource()
	//
	//err = conf.Load(src)
	////ws, err := src.Read()
	//if err != nil {
	//	log.Fatalf("Unable to load application configuration object. Err: %v\n", err)
	//	//log.Fatal(err)
	//}
	//test := conf.Map()
	////log.Printf("conf %v\n", ws.Data)
	//
	//log.Printf("conf map %v\n", test)

	audits := os.Getenv(DisableAuditRecordsEnvVarName)
	if audits == "true" {
		glDisableAuditRecords = true
	} else {
		glDisableAuditRecords = false
	}
}

func main() {

	//instantiate service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.WrapHandler(AuthWrapper),
	)

	service.Init()
	err := proto.RegisterPromotionSrvHandler(service.Server(), new(Promotion))
	if err != nil {
		log.Fatalf(glErr.SrvNoHandler(err))
	}

	//Load configuration
	loadConfig()

	myStore := service.Options().Store
	testStore(myStore)

	//Connect to DB
	conn = connectToDB()
	defer conn.Close(context.Background())

	//setup the nats broker
	mb.Br = service.Options().Broker

	// Run Service
	if err := service.Run(); err != nil {
		log.Fatalf(glErr.SrvNoStart(serviceName, err))
	}

}

func testStore(myStore store2.Store) {
	bal := myStore.Options()
	log.Printf("myStore settings %v\n", bal)
	key := "mytest"
	rec := store2.Record{
		Key:    key,
		Value:  []byte("mytest2"),
		Expiry: 2 * time.Hour,
	}

	err := myStore.Write(&rec)
	if err != nil {
		log.Printf("error writting. Error: %v", err)
	}
	rec1, err := myStore.Read(key)
	if err != nil {
		log.Printf("Uanble to read. Error: %v\n", err)
	}
	log.Printf("Read 1: %v\n", rec1)

	listLimit := store2.ListLimit(5)
	myList, err := myStore.List(listLimit)
	if err != nil {
		log.Printf("listing error %v\n", err)
	}
	log.Printf("list: %v", myList)

	err = myStore.Delete(key)
	if err != nil {
		log.Printf("delete error %v\n", err)
	}
	myList, err = myStore.List(listLimit)
	if err != nil {
		log.Printf("listing error %v\n", err)
	}
	log.Printf("list: %v", myList)
}
