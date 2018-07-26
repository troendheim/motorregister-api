package models

import (
	"motorregister-api/utils"
	"database/sql"
)

type Brand struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func Brands() *sql.Rows {
	var brandRows, err = utils.Database.Query("SELECT * FROM brand")

	if err != nil {
		panic(err.Error())
	}

	return brandRows
}