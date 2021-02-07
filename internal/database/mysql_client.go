package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lisomartinez/go-rest-example/internal/logs"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {
	db, err := sql.Open("mysql", source)
	if err != nil {
		logs.Log().Errorf("cannot create db tenant: %s", err.Error())
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to mysql!")
	}

	return &MySqlClient{db}
}
