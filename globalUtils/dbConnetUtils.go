package globalUtils

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

//PgxDBConnect: Handles connections to postgres and TimeScaleDB
type PgxDBConnect struct{}

//connectToDB: Try to connect to the DB. Return true if connection was successful, false otherwise
func (p *PgxDBConnect) ConnectToDB(databaseName string, connectionString string) (*pgx.Conn, bool, error) {
	dbConn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Printf(glErr.DbNoConnection(databaseName, err))
		return nil, false, err
	}
	return dbConn, true, nil
}

//connectToDBWithRetry: Attempts to connect to the DB every 3s for up to maxRetries in case of connection failure
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
