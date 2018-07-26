package utils

import "net/http"

func SetHTTPHeaders(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}
