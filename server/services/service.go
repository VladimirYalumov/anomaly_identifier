package services

import ai "anomaly_identifier/proto/anomaly_identifier"

type DataService interface {
	GetResult(field string, restrictiveQuery interface{}) ([]int64, error)
	InitDB(host string, user string, password string, dbName string, port string) (err error)
	MakeRestrictiveQuery(request []*ai.Limitation) (restrictiveQuery interface{})
}

type EmptyDataService struct{}

func (service *EmptyDataService) GetResult(field string, restrictiveQuery interface{}) ([]int64, error) {
	return []int64{}, nil
}

func (service *EmptyDataService) InitDB(host string, user string, password string, dbName string, port string) (err error) {
	return nil
}

func (service *EmptyDataService) MakeRestrictiveQuery(limitations []*ai.Limitation) (restrictiveQuery interface{}) {
	return
}
