package infra

import (
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Connection *sqlx.DB
}

func NewDatabase(dataSourceName string) *Database {
	connection := sqlx.MustConnect("mysql", dataSourceName)
	return &Database{connection}
}
