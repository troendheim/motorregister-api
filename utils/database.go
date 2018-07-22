package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"github.com/spf13/viper"
)

var Database *sql.DB

func OpenConnection() {
	var error error

	Database, error = sql.Open("mysql", viper.GetString("dsn"))

	if error != nil {
		log.Fatalf("Error on initializing database connection: %s", error.Error())
	}
}
