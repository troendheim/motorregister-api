package migration

import (
	"../utils"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
)

type xmlVehicle struct {
	BaseStructure struct {
		XmlVehicleDescriptionStructure struct {
			Brand string `xml:"KoeretoejMaerkeTypeNavn"`
			Model struct {
				ModelName string `xml:"KoeretoejModelTypeNavn"`
			} `xml:"Model"`
		} `xml:"KoeretoejBetegnelseStruktur"`
	} `xml:"KoeretoejOplysningGrundStruktur"`
	ZipCode      int    `xml:"AdressePostNummer"`
	LicensePlate string `xml:"RegistreringNummerNummer"`
	Status       string `xml:"KoeretoejRegistreringStatus"`
}

type importResult struct {
	Count int                               `json:"count"`
	Data  map[int]map[string]map[string]int `json:"data"`
}

func importRawData(importFile *os.File) {

	fmt.Println("Starting conversion:")

	resultDataMap := map[int]map[string]map[string]int{}
	totalCount := 0

	xmlDecoder := xml.NewDecoder(importFile)

	for {
		token, tokenErr := xmlDecoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			}
			// handle error
		}
		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "Statistik" {
				var vehicleData xmlVehicle

				if err := xmlDecoder.DecodeElement(&vehicleData, &tokenType); err != nil {
					utils.CheckError(err)
				}

				if vehicleData.ZipCode == 0 || vehicleData.Status == "Afmeldt" {
					continue
				}
				if resultDataMap[vehicleData.ZipCode] == nil {
					resultDataMap[vehicleData.ZipCode] = map[string]map[string]int{}
				}
				if resultDataMap[vehicleData.ZipCode][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand] == nil {
					resultDataMap[vehicleData.ZipCode][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand] = map[string]int{}
				}

				resultDataMap[vehicleData.ZipCode][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Model.ModelName]++

				totalCount++
				if math.Mod(float64(totalCount), 500) == 0 {
					fmt.Printf("\r... Parsed %v entries\n", totalCount)
				}
			}
		}
	}

	result := importResult{
		Count: totalCount,
		Data:  resultDataMap,
	}

	a, err := json.Marshal(result)

	utils.CheckError(err)

	ioutil.WriteFile("migration/data.json", a, 644)

	fmt.Println("Done. Wrote to migration/data.json")
}
