package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"motorregister-api/models/vehicle"
	"fmt"
	"github.com/bdwilliams/go-jsonify/jsonify"
)

// Get statistics for brand-model-zip
func ModelZipStatistics(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var models = models.GetModelCount(params.ByName("brand"), params.ByName("model"))

	fmt.Fprintf(response, "%s", jsonify.Jsonify(models))
}

// Get possible models for brand
func Models(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var models = models.Models(params.ByName("brand"))

	fmt.Fprintf(response, "%s", jsonify.Jsonify(models))
}