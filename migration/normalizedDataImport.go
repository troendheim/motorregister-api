package migration

import (
	"../utils"
	"fmt"
	"github.com/tidwall/gjson"
)

// Import statistics from data json and insert into DB
func importNormalizedDataToDB(importFileJson gjson.Result) {
	totalCount := importFileJson.Get("count").Int()
	fmt.Printf("\nStarting import of statistics based on '%v' vehicles\n", totalCount)

	importFileJson.Get("data").ForEach(func(zipCode, zipCodeData gjson.Result) bool {
		fmt.Printf(".. Importing zip code '%s' ", zipCode)

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

		fmt.Printf("-- %v%% of all vehicles loaded \n", float32(currentDoneCount)*float32(100)/float32(totalCount))

		return true
	})
}
