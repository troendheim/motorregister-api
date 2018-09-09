package migration

import (
	"../models/location"
	"../utils"
	"encoding/json"
	"fmt"
	"net/http"
)

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

func importZipCodeData() {
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
