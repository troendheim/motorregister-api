package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"motorregister-api/utils"
	"motorregister-api/config"
	"github.com/spf13/viper"
)

// Initialize application
func main() {
	// Start up the cache
	utils.StartCache()

	// Read configuration
	config.PrepareConfig()

	// Get port from app.yml
	serverPort := viper.GetInt("port")
	fmt.Println("Starting server on port", serverPort)

	// Build routes
	router := buildRouter()

	// Get DB connection ready
	utils.OpenConnection()

	// Start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), router))
}
