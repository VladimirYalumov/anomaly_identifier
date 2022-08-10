package pgsql_service

import (
	ai "anomaly_identifier/proto/anomaly_identifier"
	"anomaly_identifier/server/services/pgsql_service/db"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type PgSQLDataService struct {
	PgClient db.PgSQLClient
	Table    string
}

func (service *PgSQLDataService) GetResult(field string, restrictiveQuery interface{}) ([]int64, error) {
	sqlQuery, ok := restrictiveQuery.(string)
	if !ok {
		return nil, errors.New("grpc: unidentified limitations in service: PgSQLDataService")
	}

	return nil, nil
}

func (service *PgSQLDataService) InitDB(host string, user string, password string, dbName string, port string) (err error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	service.PgClient.Database, err = gorm.Open("postgres", dbinfo)
	return err
}

func (service *PgSQLDataService) MakeRestrictiveQuery(limitations []*ai.Limitation) (restrictiveQuery interface{}) {
	limitationsCount := len(limitations) - 1
	for iterator, limitation := range limitations {
		restrictiveQuery = fmt.Sprintf(
			"%s %s %s %v",
			restrictiveQuery,
			limitation.GetField(),
			limitation.GetComparisonSign(),
			limitation.GetValue(),
		)
		if iterator != limitationsCount {
			restrictiveQuery = fmt.Sprintf("%s and ", restrictiveQuery)
		}
	}
	return
}
