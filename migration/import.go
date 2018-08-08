package migration

import (
	"io/ioutil"
	"github.com/tidwall/gjson"
	"fmt"
	"../utils"
	"path/filepath"
	"os"
	"strings"
	"../models/location"
	"net/http"
	"encoding/json"
)

func Import(dataImportFileLocation *string) {
	importFileContents, err := ioutil.ReadFile(*dataImportFileLocation)

	if err != nil {
		panic(err.Error())
	}

	applyPatches()
	doImportStatistics(gjson.Parse(string(importFileContents)))
	doImportZipcodeData()
}

// Apply patches from ./db/ folder
func applyPatches() {
	fmt.Println("Applying patches")

	filepath.Walk("migration/db/", func(path string, info os.FileInfo, err error) error {
		if ! info.IsDir() {
			var patchFile, patchFileReadError = ioutil.ReadFile(path)
			if patchFileReadError != nil {
				panic(patchFileReadError.Error())
			}

			fmt.Printf("\n.. %s\n", info.Name())

			var fullQueryContent = string(patchFile)

			var queries = strings.Split(fullQueryContent, ";")
			for index, query := range queries {
				if query == "\n" {
					continue
				}
				fmt.Printf(".... Query %v\n", index+1)
				utils.Database.Exec(query)
			}
		}

		return nil
	})
}

// Import statistics from data json and insert into DB
func doImportStatistics(importFileJson gjson.Result) {
	totalCount := importFileJson.Get("count").Int()
	fmt.Printf("\nStarting import of statistics based on '%v' vehicles\n", totalCount)

	importFileJson.Get("data").ForEach(func(zipCode, zipCodeData gjson.Result) bool {
		fmt.Printf(".. Importing zip code '%s' ", zipCode)

		// Insert zip if not exists
		utils.Database.Exec(`
				INSERT INTO zip_code (zip_code)
				VALUES (?)
				ON DUPLICATE KEY UPDATE zip_code = zip_code
			`, zipCode.String())

		var zipCodeId int
		utils.Database.QueryRow(`
				SELECT	id
				FROM	zip_code
				WHERE	zip_code = ?
			`, zipCode.String()).Scan(&zipCodeId)

		zipCodeData.ForEach(func(makeName, models gjson.Result) bool {
			// Insert make if not exists
			utils.Database.Exec(`
				INSERT INTO brand (name)
				VALUES (?)
				ON DUPLICATE KEY UPDATE name = name
			`, makeName.String())

			var brandId int
			utils.Database.QueryRow(`
				SELECT	id
				FROM	brand
				WHERE	name = ?
			`, makeName.String()).Scan(&brandId)

			models.ForEach(func(modelName, modelCount gjson.Result) bool {
				// Insert model
				utils.Database.Exec(`
					INSERT IGNORE INTO model (name, brand_id)
					VALUES (?, ?)
					ON DUPLICATE KEY UPDATE name = name
				`, modelName.String(), brandId)

				var modelId int
				utils.Database.QueryRow(`
					SELECT	id
					FROM	model
					WHERE	brand_id = ? AND name = ?
				`, brandId, modelName.String()).Scan(&modelId)

				// Add stats
				utils.Database.Exec(`
					INSERT INTO model_2_zip_count
					VALUES (?, ?, ?)
				`, zipCodeId, modelId, modelCount.Int())

				return true
			})
			return true
		})

		var currentDoneCount int64
		utils.Database.QueryRow(`
			SELECT	SUM(total_count) as currentDoneCount
			FROM	model_2_zip_count
		`).Scan(&currentDoneCount)

		fmt.Printf("-- %v%% of all vehicles loaded \n", float32(currentDoneCount) * float32(100) / float32(totalCount))

		return true
	})
}

type geocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
}

func doImportZipcodeData() {
	var zipCodes = models.GetAllZipCodes()

	fmt.Println("Importing locations for zip codes")

	for zipCodes.Next() {
		var zipCode int
		zipCodes.Scan(&zipCode)

		fmt.Printf(".. %v\n", zipCode)

		var response, err = http.Get(fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=%v,Denmark", zipCode))
		if err != nil {
			panic(err.Error())
		}

		responseBody := &geocodeResponse{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if len(responseBody.Results) < 1 {
			fmt.Printf(".... Could not map to a location. Damn you google..\n")
			continue
		}

		location := responseBody.Results[0].Geometry.Location

		utils.Database.Exec(`
			UPDATE	zip_code
			SET		latitude = ?, longtitude = ?
			WHERE	zip_code = ?
		`, location.Lat, location.Lng, zipCode)
	}
}
