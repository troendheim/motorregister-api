package main

import (
	"./config"
	"./migration"
	"./utils"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// Initialize application
func main() {
	// Init flags
	normalizedDataImportFile := flag.String("normalizedDataImportFile", "", "Set -normalizedDataImportFile flag to start migration based on json file")
	rawDataImportFile := flag.String("rawDataImportFile", "", "Set -rawDataImportFile flag to start migration based on raw XML")
	importZipDataFlag := flag.Bool("importZipData", true, "Set -importZipData flag to start import of zip code data")
	flag.Parse()

	// Read configuration
	config.PrepareConfig()

	// Get DB connection ready
	utils.OpenConnection()

	if *normalizedDataImportFile != "" {

		migration.Import("normalized", normalizedDataImportFile)

	} else if *rawDataImportFile != "" {

		migration.Import("raw", rawDataImportFile)

	} else if *importZipDataFlag == true {

		migration.Import("zip_code_data", normalizedDataImportFile) // normalizedDataImportFile <- Dummy arg

	} else {

		// Start up the cache
		utils.StartCache()

		// Build routes
		router := BuildRouter()

		// Get port from app.yml
		serverPort := viper.GetInt("port")
		fmt.Println("Starting server on port", serverPort)

		// Start server
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), router))
	}
}
