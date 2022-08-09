package services

type DataService interface {
	GetResult(field string, restrictiveQuery string) ([]int32, error)
	InitDB(host string, user string, password string, dbName string, port string) (err error)
}
