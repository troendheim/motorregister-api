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
func Models(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	utils.SetHTTPHeaders(responseWriter)

	var brandName = params.ByName("brand")

	var cacheKey = fmt.Sprintf("cached_response_brands_%s", brandName);

	var response, cacheError = utils.GetFromCache(cacheKey)
	if cacheError != nil {
		var models = models.Models(brandName)
		response = utils.ConvertStringSliceToString(jsonify.Jsonify(models))

		utils.SetCache(cacheKey, response)
	}


	fmt.Fprint(responseWriter, response)
}

func Brands(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	utils.SetHTTPHeaders(responseWriter)

	var cacheKey = "cached_response_brands";

	var response, cacheError = utils.GetFromCache(cacheKey)
	if cacheError != nil {
		var brands = models.Brands()
		response = utils.ConvertStringSliceToString(jsonify.Jsonify(brands))

		utils.SetCache(cacheKey, response)
	}

	fmt.Fprint(responseWriter, response)
}