package migration

import (
	"../utils"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
)

func Import(importType string, dataImportFileLocation *string) {
	switch importType {
	case "normalized":
		importFileContents, err := ioutil.ReadFile(*dataImportFileLocation)
		utils.CheckError(err)

		applyPatches()
		importZipCodeData()
		importNormalizedDataToDB(gjson.Parse(string(importFileContents)))
	case "raw":
		importFileReader, err := os.Open(*dataImportFileLocation)
		utils.CheckError(err)

		importRawData(importFileReader)

	case "zip_code_data":

		importZipCodeData()
	}

}
