package main

import (
	"github.com/julienschmidt/httprouter"
	"motorregister-api/controllers"
	"github.com/damnever/cc"
)

func buildRouter(config *cc.Config) *httprouter.Router {
	router := httprouter.New()
	router.GET("/models/:brand/:model", controllers.GetModelsAction)

	return router
}
