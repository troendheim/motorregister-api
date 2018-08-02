package models

import (
	"motorregister-api/utils"
	"database/sql"
)

type ZipCode struct {
	Id         int `json:"id" db:"id"`
	ZipCode    int `json:"zip_code" db:"zip_code"`
	Name       string
	Latitude   float64
	Longtitude float64
}

func GetAllZipCodes() sql.Rows {
	var rows, err = utils.Database.Query("SELECT zip_code FROM zip_code")
	if err != nil {
		panic(err.Error())
	}

	return *rows
}