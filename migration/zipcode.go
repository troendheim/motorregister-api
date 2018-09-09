package migration

import (
	"../utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type geocodeResponse struct {
	ZipCode      string    `json:"nr"`
	Name         string    `json:"navn"`
	VisualCenter []float64 `json:"visueltcenter"`
}
type geocodeResponseArray []geocodeResponse

func importZipCodeData() {
	fmt.Println("Importing locations for zip codes")

	// Fetch from DAWA
	var zipCodeDataFromDawaResponse, err = http.Get("https://dawa.aws.dk/postnumre")
	utils.CheckError(err)
	zipCodeDataFromDawa := &geocodeResponseArray{}
	json.NewDecoder(zipCodeDataFromDawaResponse.Body).Decode(&zipCodeDataFromDawa)

	for _, zipCodeData := range *zipCodeDataFromDawa {
		utils.Database.Exec(`
			INSERT INTO	zip_code
			(zip_code, latitude, longtitude, name) VALUES
			(?, ?, ?, ?)
		`, zipCodeData.ZipCode, zipCodeData.VisualCenter[1], zipCodeData.VisualCenter[0], zipCodeData.Name)

		fmt.Printf(".. %v\n", zipCodeData.ZipCode)
	}
}
