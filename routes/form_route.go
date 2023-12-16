package routes

import (
	"manifest-craft/controllers"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

func FormRoute(router *gin.Engine, store storage.Storage) {
	formHandler := controllers.NewFormsHandler(store)

	router.GET("/form/:name", formHandler.Get)
}
