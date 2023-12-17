package routes

import (
	"manifest-craft/controllers"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

func ManifestRoute(router *gin.Engine, store storage.Storage) {
	manifestController := controllers.NewManifestController(store)

	router.GET("/form/:name", manifestController.Get)
	router.POST("/form/:name", manifestController.Send)
}
