package routes

import (
	"manifest-craft/controllers"
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

func ComponentRoute(router *gin.Engine, store storage.Storage, services *services.Services) {
	componentController := controllers.NewComponentController(store, services)

	router.GET("/values/:source", componentController.GetValues)
}
