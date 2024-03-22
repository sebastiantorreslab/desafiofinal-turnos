package store

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

type SqlStore struct {
	StorageDB *sql.DB
}

func init() {

	dataSource := "user-db:pass-db@tcp(localhost:3360)/clinic-db"

	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	StorageDB.SetConnMaxLifetime(time.Minute * 3)
	StorageDB.SetMaxOpenConns(10)
	StorageDB.SetMaxIdleConns(10)

	log.Println("database initialized")

}
