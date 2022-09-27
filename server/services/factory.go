package services

import (
	ai "anomaly_identifier/proto/anomaly_identifier"
	"anomaly_identifier/server/services/pgsql_service"
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var configDbName = map[ai.Service]string{
	ai.Service_PGSQL: "pgsql",
}

func GetService(request *ai.AnomalyIdentifierRequest) (DataService, error) {
	switch request.Service {
	case ai.Service_PGSQL:
		service := &pgsql_service.PgSQLDataService{}
		err := service.InitDB(
			connectionParams.Host,
			connectionParams.User,
			connectionParams.Password,
			connectionParams.dbName,
			connectionParams.Port,
		)
		if err != nil {
			return &EmptyDataService{}, err
		}
		return service, nil
	default:
		return &EmptyDataService{}, errors.New("grpc: unidentified service")
	}
}

type PGSQLConnectionConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	dbName   string `yaml:"db_name"`
	Port     string `yaml:"port"`
}

func getConnection(serviceId int32) (configurations DbConnectionConfig) {
	yfile, parseErr := ioutil.ReadFile("db_connection.yml")
	if parseErr != nil {
		return
	}
	parseErr = yaml.Unmarshal(yfile, &configurations)
	if parseErr != nil {
		panicIfNeed(parseErr, "")
	}
}
