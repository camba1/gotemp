package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/micro/go-micro/v2"
	"goTemp/globalUtils"
	"goTemp/globalerrors"
	"log"
	"os"
)

//glErr: Holds the service global errors that are shared cross services
var glErr globalerrors.SrvError

const (
	//serviceName: Name of this service and which will be used for service discovery and registration
	serviceName = "goTemp.api.audit"
)

const (
	//dbName: Name of the DB hosting the data
	dbName = "appuser"
	//dbConStrEnvVarName: Name of Environment variable that contains connection string to DB
	dbConStrEnvVarName = "DB_CONNECT"
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

//Authentication middleware
//func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
//	return func(ctx context.Context, req server.Request, resp interface{}) error {
//		//User login is excepted from authentication
//		if req.Endpoint() == "UserSrv.Auth" {
//			return fn(ctx, req, resp)
//		}
//		meta, ok := metadata.FromContext(ctx)
//		if !ok {
//			return fmt.Errorf(glErr.AuthNoMetaData(req.Endpoint()))
//		}
//
//		auth,ok := meta["Authorization"]
//		if !ok {
//			return fmt.Errorf(glErr.AuthNilToken())
//		}
//		authSplit := strings.SplitAfter(auth, " ")
//			if len(authSplit) != 2 {
//		return fmt.Errorf(glErr.AuthNilToken())
//}
//		token := authSplit[1]

//		//token := meta["Token"]
//		log.Printf("endpoint: %v", req.Endpoint())
//
//		var u User
//		outToken := &pb.Token{}
//		err := u.ValidateToken(ctx, &pb.Token{Token: token}, outToken)
//		//authClient := userService.NewAuthClient("shippy.user", srv.Client())
//		//_, err := authClient.ValidateToken(ctx, &userService.Token{
//		//	Token: token,
//		//})
//		if err != nil {
//			return err
//		}
//		if outToken.Valid != true {
//			return fmt.Errorf(glErr.AuthInvalidToken())
//		}
//		return fn(ctx, req, resp)
//	}
//}
func connectToDB() *pgx.Conn {
	var pgxConnect globalUtils.PgxDBConnect
	dbConn, err := pgxConnect.ConnectToDBWithRetry(dbName, getDBConnString())
	if err != nil {
		log.Fatalf(glErr.DbNoConnection(dbName, err))
	}
	return dbConn
}

func main() {

	// setup metrics collector
	//metricsWrapper := newMetricsWrapper()
	//metricsSubscriberWrapper := newMetricsSubscriberWrapper()
	//instantiate service
	service := micro.NewService(
		micro.Name(serviceName),
		//micro.WrapHandler(AuthWrapper),
		//micro.WrapHandler(metricsWrapper),
		//micro.WrapSubscriber(metricsSubscriberWrapper),
	)

	service.Init()
	//err := pb.RegisterUserSrvHandler(service.Server(), new(User))
	//if err != nil {
	//	log.Fatalf(glErr.SrvNoHandler(err))
	//}

	//Connect to DB
	conn = connectToDB()
	defer conn.Close(context.Background())

	//setup the nats broker
	mb.Br = service.Options().Broker
	defer mb.Br.Disconnect()

	//mb.SubscribeHandleWrapper = globalMonitoring.NewMetricsSubscriberWrapper(
	//	globalMonitoring.ServiceName(serviceName),
	//	globalMonitoring.ServiceID(serviceId),
	//	globalMonitoring.ServiceVersion("latest"),
	//	globalMonitoring.ServiceMetricsPrefix(serviceMetricsPrefix),
	//	globalMonitoring.ServiceMetricsLabelPrefix(serviceMetricsPrefix),
	//)

	// setup metrics collector for broker messages
	mb.SubscribeHandleWrapper = newMetricsSubscriberWrapper()

	//topic := "test"
	//queueName := "test"
	//_ = mb.subToMsg( getMsg, topic,queueName)
	//userToSend := &pb.User{Id: 1234, Email: "werewq", ValidFrom: ptypes.TimestampNow()}
	//header := map[string]string{"test": "Test"}
	//_ = mb.sendMsg( userToSend,header, topic)

	//setup broker subscription
	var aud AuditSrv
	err := aud.SubsToBrokerInsertMsg()
	if err != nil {
		log.Printf("Error subscribing to message: Error: %v\n", err)
	}

	// Initialize http server for metrics export
	go runHttp()

	// Run Service
	err = service.Run()
	if err != nil {
		log.Fatalf(glErr.SrvNoStart(serviceName, err))
	}
}
