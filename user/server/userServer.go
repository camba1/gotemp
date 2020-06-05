package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
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

//Authentication middleware
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		//User login is excepted from authentication
		if req.Endpoint() == "UserSrv.Auth" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return fmt.Errorf(glErr.AuthNoMetaData(req.Endpoint()))
		}

		token := meta["Token"]
		log.Printf("endpoint: %v", req.Endpoint())

		var u User
		outToken := &pb.Token{}
		err := u.ValidateToken(ctx, &pb.Token{Token: token}, outToken)
		//authClient := userService.NewAuthClient("shippy.user", srv.Client())
		//_, err := authClient.ValidateToken(ctx, &userService.Token{
		//	Token: token,
		//})
		if err != nil {
			return err
		}
		if outToken.Valid != true {
			return fmt.Errorf(glErr.AuthInvalidToken())
		}
		return fn(ctx, req, resp)
	}
}

func main() {

	//instantiate service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.WrapHandler(AuthWrapper),
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

	//setup the nats broker

	mb.br = service.Options().Broker
	//topic := "test"
	//queueName := "test"
	//_ = mb.subToMsg( getMsg, topic,queueName)
	//userToSend := &pb.User{Id: 1234, Email: "werewq", ValidFrom: ptypes.TimestampNow()}
	//header := map[string]string{"test": "Test"}
	//_ = mb.sendMsg( userToSend,header, topic)

	// Run Service
	err = service.Run()
	if err != nil {
		log.Fatalf(glErr.SrvNoStart(serviceName, err))
	}
}
