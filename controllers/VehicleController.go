package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"motorregister-api/models/vehicle"
	"github.com/bdwilliams/go-jsonify/jsonify"
	"fmt"
)

func ModelZipStatistics(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var models = models.GetModelCount(params.ByName("brand"), params.ByName("model"))

	fmt.Fprintf(response, "%s", jsonify.Jsonify(models))
}
