package routes

import (
	"manifest-craft/controllers"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

func ManifestRoute(router *gin.Engine, store storage.Storage) {
	manifestHandler := controllers.NewManifestHandler(store)

	router.GET("/form/:name", manifestHandler.Get)
}
