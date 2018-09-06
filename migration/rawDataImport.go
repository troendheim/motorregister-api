package migration

import (
	"../utils"
	"encoding/xml"
	"fmt"
	"io"
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
	ZipCode int `xml:"AdressePostNummer"`
	LicensePlate string `xml:"RegistreringNummerNummer"`
}

type result struct {
	Count int
	Data struct {
		ZipCode map[int] struct {
			Brand map[string] struct {
				Model map[string] struct {
					count int
				}
			}
		}
	}
}

func importRawData(importFile *os.File) {

	fmt.Println("Starting conversion ..")

	var resultData result

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
					// handle error
				}

				if vehicleData.ZipCode == 0 {
					continue
				}

				resultData.Count++

				resultData.Data.ZipCode[vehicleData.ZipCode] = map[string]{map[string] = 1}


				//[vehicleData.ZipCode][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand][vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Model.ModelName] = 1

				fmt.Println(vehicleData.LicensePlate, vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Brand, vehicleData.BaseStructure.XmlVehicleDescriptionStructure.Model.ModelName, vehicleData.ZipCode)

				// do something with vehicleData
			}
		}
	}

}
