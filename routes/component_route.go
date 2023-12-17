package routes

import (
	"manifest-craft/controllers"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

func ComponentRoute(router *gin.Engine, store storage.Storage) {
	componentHandler := controllers.NewComponentHandler(store)

	router.GET("/values/:source", componentHandler.GetValues)
}
