package globalUtils

import (
	"context"
	adb "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

// PgxDBConnect handles connections to postgres and TimeScaleDB
type PgxDBConnect struct{}

// ConnectToDB tries to connect to the DB. Return true if connection was successful, false otherwise
func (p *PgxDBConnect) ConnectToDB(databaseName string, connectionString string) (*pgx.Conn, bool, error) {
	dbConn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Printf(glErr.DbNoConnection(databaseName, err))
		return nil, false, err
	}
	return dbConn, true, nil
}

// ConnectToDBWithRetry attempts to connect to the DB every 3s for up to maxRetries in case of connection failure
func (p *PgxDBConnect) ConnectToDBWithRetry(databaseName string, connectionString string) (*pgx.Conn, error) {
	maxRetries := 5
	var dbConn *pgx.Conn
	var connected bool
	var err error
	for i := 1; i <= maxRetries; i++ {
		dbConn, connected, err = p.ConnectToDB(databaseName, connectionString)
		if !connected {
			if i >= maxRetries {
				log.Printf(glErr.DbNoConnection(databaseName, err))
				return nil, err
			} else {
				log.Printf(glErr.DbConnectRetry(i, err))
				time.Sleep(3 * time.Second)
			}
		} else {
			break
		}
	}
	return dbConn, nil
}

// DbConnParams defines generic connection parameters to connect to the DB
type DbConnParams struct {
	Address  string
	Username string
	Password string
}

// ArangoConnect handles connections to ArangoDB
type ArangoConnect struct{}

// ConnectToDB tries to connect to the DB. Return true if connection was successful, false otherwise
func (ac *ArangoConnect) ConnectToDB(databaseName string, connParams *DbConnParams) (adb.Database, bool, error) {

	//  Create an HTTP connection to the database
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + connParams.Address},
	})
	if err != nil {
		return nil, false, nil
	}
	//  Create a client
	cli, err := adb.NewClient(adb.ClientConfig{
		Connection: conn,
		// Authentication: adb.JWTAuthentication('ere','ere'),
		Authentication: adb.BasicAuthentication(connParams.Username, connParams.Password),
	})
	if err != nil {
		return nil, false, err
	}

	ctx := context.Background()
	db, err := cli.Database(ctx, databaseName)
	if err != nil {
		return nil, false, err
	}

	return db, true, nil
}

// ConnectToDBWithRetry attempts to connect to the DB every 3s for up to maxRetries in case of connection failure
func (ac *ArangoConnect) ConnectToDBWithRetry(databaseName string, connParams *DbConnParams) (adb.Database, error) {
	maxRetries := 5
	var db adb.Database
	var connected bool
	var err error
	for i := 1; i <= maxRetries; i++ {
		db, connected, err = ac.ConnectToDB(databaseName, connParams)
		if !connected {
			if i >= maxRetries {
				log.Printf(glErr.DbNoConnection(databaseName, err))
				return nil, err
			} else {
				log.Printf(glErr.DbConnectRetry(i, err))
				time.Sleep(3 * time.Second)
			}
		} else {
			break
		}
	}
	log.Println("Connected to DB")
	return db, nil
}
