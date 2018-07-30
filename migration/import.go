package migration

import (
	"io/ioutil"
	"github.com/tidwall/gjson"
	"fmt"
	"motorregister-api/utils"
	"path/filepath"
	"os"
	"strings"
)

func Import(dataImportFileLocation *string) {
	importFileContents, err := ioutil.ReadFile(*dataImportFileLocation)

	if err != nil {
		panic(err.Error())
	}

	applyPatches()
	doImport(gjson.Parse(string(importFileContents)))
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
				fmt.Printf(".... Query %v\n", index + 1)
				utils.Database.Exec(query)
			}
		}


		return nil
	})
}

// Import statistics from data json and insert into DB
func doImport(importFileJson gjson.Result) {
	totalCount := importFileJson.Get("count").Int()
	fmt.Printf("\nStarting import of statistcs based on '%v' vehicles\n", totalCount)

	importFileJson.Get("data").ForEach(func(zipCode, zipCodeData gjson.Result) bool {
		fmt.Printf(".. Importing zip code '%s' \n", zipCode)
		zipCodeData.ForEach(func(makeName, models gjson.Result) bool {
			// Insert make if not exists
			utils.Database.Exec(`
				INSERT INTO brand (name)
				VALUES (?)
				ON DUPLICATE KEY UPDATE name = name
			`, makeName.String())

			models.ForEach(func(modelName, modelCount gjson.Result) bool {
				// Insert model
				utils.Database.Exec(`
					INSERT IGNORE INTO model (name, brand_id)
					VALUES (
						?,
						( SELECT id
						  FROM brand
                          WHERE name = ?
					)
					ON DUPLICATE KEY UPDATE name = name
				`, modelName.String(), makeName.String())

				// Add stats
				utils.Database.Exec(`
					INSERT INTO model_2_zip_count
					VALUES (
						# Zip code
						(SELECT id FROM zip_code WHERE zip_code = ?),

						# Model id
						(SELECT id FROM model WHERE name = ?),

						# Count
						?
					)
				`, zipCode.Int(), modelName.String(), totalCount)

				return true
			})
			return true
		})

		return true
	})
}
