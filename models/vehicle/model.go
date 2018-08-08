package models

import (
	"../../utils"
	"database/sql"
)

type Model struct {
	Id      int		`db:"id"`
	Name    string	`db:"name"`
	BrandId int		`db:"brand_id"`
}

// Get model count for brand-model
func GetModelCount(brandName string, modelName string) *sql.Rows {
	var preparedQuery, prepareError = utils.Database.Prepare(`
SELECT zip_code as zipCode, latitude, longtitude, brand.name as brandName, model.name as modelName, total_count as totalCount
FROM brand
JOIN model
  ON model.brand_id = brand.id
  AND model.name = ?
JOIN model_2_zip_count
  ON model_2_zip_count.model_id = model.id
JOIN zip_code
  ON model_2_zip_count.zip_code_id = zip_code.id 
WHERE brand.name = ?
`)
	if prepareError != nil {
		panic(prepareError.Error())
	}

	var rows, queryError = preparedQuery.Query(modelName, brandName)
	if queryError != nil {
		panic(queryError.Error())
	}

	return rows
}

// Get possible models for a given brand
func Models(brandName string) *sql.Rows {
	var preparedQuery, prepareError = utils.Database.Prepare(`
SELECT model.*
FROM brand
JOIN model
  ON model.brand_id = brand.id 
WHERE brand.name = ?
`)
	if prepareError != nil {
		panic(prepareError.Error())
	}

	var rows, queryError = preparedQuery.Query(brandName)
	if queryError != nil {
		panic(queryError.Error())
	}

	return rows
}
