package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"motorregister-api/utils"
	"motorregister-api/config"
)

// Initialize application
func main() {
	// Read configuration
	config := config.GetConfig()

	// Get port from app.yml
	serverPort := config.IntOr("port", 8999)
	fmt.Println("Starting server on port", serverPort)

	// Build routes
	router := buildRouter(config)

	// Get DB connection ready
	utils.OpenConnection(config)

	// Start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), router))
}
