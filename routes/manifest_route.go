package routes

import (
	"manifest-craft/controllers"
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

func ManifestRoute(router *gin.Engine, store storage.Storage, services services.Services) {
	manifestController := controllers.NewManifestController(store, services)

	router.GET("/form/:name", manifestController.Get)
	router.POST("/form/:name", manifestController.Send)
}
