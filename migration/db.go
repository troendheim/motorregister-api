package migration

import (
	"../utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Apply patches from ./db/ folder
func applyPatches() {
	fmt.Println("Applying patches")

	filepath.Walk("migration/db/", func(path string, info os.FileInfo, err error) error {
		if ! info.IsDir() {
			var patchFile, patchFileReadError = ioutil.ReadFile(path)
			if patchFileReadError != nil {
				panic(patchFileReadError.Error())
			}

			fmt.Printf("\n.. %s\n", info.Name())

			var fullQueryContent = string(patchFile)

			var queries = strings.Split(fullQueryContent, ";")
			for index, query := range queries {
				if query == "\n" {
					continue
				}
				fmt.Printf(".... Query %v\n", index+1)
				utils.Database.Exec(query)
			}
		}

		return nil
	})
}
