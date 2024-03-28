package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {

	dataSource := "user-db:pass-db@tcp(localhost:3360)/clinic-db"

	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db, nil
}