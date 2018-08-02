package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"motorregister-api/utils"
	"motorregister-api/config"
	"github.com/spf13/viper"
	"flag"
	"motorregister-api/migration"
)

// Initialize application
func main() {
	// Init flags
	dataImportFile := flag.String("dataImportFile", "", "Set -dataImportFile flag to start migration based on json file")
	flag.Parse()

	// Start up the cache
	utils.StartCache()

	// Read configuration
	config.PrepareConfig()

	// Build routes
	router := buildRouter()

	// Get DB connection ready
	utils.OpenConnection()

	if *dataImportFile != "" {
		migration.Import(dataImportFile)
	} else {
		// Start server
		// Get port from app.yml
		serverPort := viper.GetInt("port")
		fmt.Println("Starting server on port", serverPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), router))
	}
}
