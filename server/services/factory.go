package services

import (
	ai "anomaly_identifier/proto/anomaly_identifier"
	"anomaly_identifier/server/services/pgsql_service"
	"errors"
)

var ConnectionSettings = [5]string{}

func GetService(request *ai.AnomalyIdentifierRequest) (DataService, error) {
	switch request.Service {
	case ai.Service_PGSQL:
		setConnectionSettings("127.0.0.1", "user", "qwerty", "deewave", "")
		service := &pgsql_service.PgSQLDataService{}
		return service, nil
	default:
		return &EmptyDataService{}, errors.New("grpc: unidentified service")
	}
}

func setConnectionSettings(host string, user string, password string, dbName string, port string) {
	ConnectionSettings[0] = host
	ConnectionSettings[1] = user
	ConnectionSettings[2] = password
	ConnectionSettings[3] = dbName
	ConnectionSettings[4] = port
}
