package migration

import (
	"../utils"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"os"
)

type xmlVehicleBaseStructure struct {
}

type xmlVehicleDescriptionStructure struct {
}

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
}

type result struct {
	Count int
	Data  struct {
		ZipCode map[int]struct {
			Brand map[string]struct {
				Model map[string]struct {
					Count int
				}
			}
		}
	}
}

func importRawData(importFile *os.File) {

	fmt.Println("Starting conversion:")

	resultData := map[int]map[string]map[string]int{}
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

				if vehicleData.ZipCode == 0 {
					continue
				}
				if resultData[vehicleData.ZipCode] == nil {
					resultData[vehicleData.ZipCode] = map[string]map[string]int{}
				}
				if resultData[vehicleData.ZipCode][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand] == nil {
					resultData[vehicleData.ZipCode][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand] = map[string]int{}
				}

				resultData[vehicleData.ZipCode][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Model.ModelName]++

				totalCount++
				if math.Mod(float64(totalCount), 125) == 0 {
					fmt.Printf("\r... Parsed %v entries", totalCount)
				}
			}
		}
	}

	// @TODO: Convert resultData to JSON based and place file in migration

}
