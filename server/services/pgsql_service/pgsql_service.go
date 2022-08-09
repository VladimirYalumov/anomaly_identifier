package pgsql_service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"grpc_anomaly_identifier/server/services/pgsql_service/db"
)

type PgSQLDataService struct {
	PgClient db.PgSQLClient
}

func (service *PgSQLDataService) GetResult(table string, field string, restrictiveQuery string) ([]int32, error) {
	return nil, nil
}

func (service *PgSQLDataService) InitDB(host string, user string, password string, dbName string, port string) (err error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	service.PgClient.Database, err = gorm.Open("postgres", dbinfo)
	return err
}
