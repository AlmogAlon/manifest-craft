package routes

import (
	"manifest-craft/controllers"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

func ComponentRoute(router *gin.Engine, store storage.Storage) {
	componentController := controllers.NewComponentController(store)

	router.GET("/values/:source", componentController.GetValues)
}
