package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"github.com/spf13/viper"
	"time"
)

var Database *sql.DB

func OpenConnection() {
	var error error

	Database, error = sql.Open("mysql", viper.GetString("dsn"))

	Database.SetConnMaxLifetime(time.Minute * 2)
	Database.SetMaxOpenConns(20)
	Database.SetMaxIdleConns(5)

	if error != nil {
		log.Fatalf("Error on initializing database connection: %s", error.Error())
	}
}
