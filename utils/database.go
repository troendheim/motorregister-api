package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"github.com/damnever/cc"
)

var Database *sql.DB

func OpenConnection(config *cc.Config) {
	var error error

	Database, error = sql.Open("mysql", 	config.String("dsn"))

	if error != nil {
		log.Fatalf("Error on initializing database connection: %s", error.Error())
	}
}
