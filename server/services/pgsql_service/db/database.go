package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PgSQLClient struct {
	Database *gorm.DB
}

func (client *PgSQLClient) GetValues(table string, field string, restrictiveQuery string) {
	result := map[string]interface{}{}
	client.Database.Table("users").Take(&result, restrictiveQuery)
}
