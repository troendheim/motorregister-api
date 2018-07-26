package main

import (
	"github.com/julienschmidt/httprouter"
	"motorregister-api/controllers"
)

func buildRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/statistics/model/:brand/:model", controllers.ModelZipStatistics)
	router.GET("/models/:brand", controllers.Models)
	router.GET("/brands", controllers.Brands)

	return router
}
