package utils

import "strings"

func ConvertStringSliceToByteSlice(string []string) (byteSlice []byte) {
	var joinedStrings = strings.Join(string, "")
	byteSlice = []byte(joinedStrings)

	return byteSlice
}

func ConvertStringSliceToString(string []string) (string) {
	return strings.Join(string, "")
}

func ConvertStringToByteSlice(str string) ([]byte) {
	return []byte(str)
}