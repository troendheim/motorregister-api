package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"motorregister-api/models/vehicle"
	"fmt"
	"github.com/bdwilliams/go-jsonify/jsonify"
	"motorregister-api/utils"
)

// Get statistics for brand-model-zip
func ModelZipStatistics(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	utils.SetHTTPHeaders(response)

	var models = models.GetModelCount(params.ByName("brand"), params.ByName("model"))
	fmt.Fprint(response, jsonify.Jsonify(models))
}

// Get possible models for brand
func Models(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	utils.SetHTTPHeaders(response)

	var models = models.Models(params.ByName("brand"))
	fmt.Fprint(response, jsonify.Jsonify(models))
}

func Brands(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	utils.SetHTTPHeaders(response)

	var brands = models.Brands()
	fmt.Fprint(response, jsonify.Jsonify(brands))
}