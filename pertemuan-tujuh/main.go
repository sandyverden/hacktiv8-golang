package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username   string = "root"
	password   string = ""
	hostname   string = "127.0.0.1"
	dbName     string = "kampus_merdeka"
	connection string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, hostname, dbName)
)

type Student struct {
	id    string
	name  string
	age   int
	grade int
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	SqlQuery()
}
