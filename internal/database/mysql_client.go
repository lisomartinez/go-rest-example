package database

import (
	"database/sql"
	"fmt"
)

type MySqlClient struct {
}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)
	if err != nil {
		_ = fmt.Errorf("cannot create db tenant: %s", err.Error())
		panic("cannot create db tenant")
	}

	return db
}
