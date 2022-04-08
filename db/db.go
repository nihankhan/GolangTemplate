package db

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	username = "root"
	password = "nihankhan"
	hostname = "127.0.0.1:3306"
	dbName   = "Nihan"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

var (
	db  *sql.DB
	err  error
)

func Connect() (db *sql.DB) {
	db, err = sql.Open(dbDriver, dsn(""))

	if err != nil {
		panic(err)
	}

	return db
}