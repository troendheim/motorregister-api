package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"motorregister-api/utils"
	"motorregister-api/config"
)

func main() {
	fmt.Println("Starting server on port 8999")

	config := config.GetConfig()

	router := buildRouter(config)

	utils.OpenConnection(config)

	log.Fatal(http.ListenAndServe(":8999", router))
}
