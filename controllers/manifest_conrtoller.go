package controllers

import (
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

type ManifestController struct {
	store    storage.Storage
	services services.Services
}

func NewManifestController(storage storage.Storage, services services.Services) *ManifestController {
	return &ManifestController{
		store:    storage,
		services: services,
	}
}

func (c *ManifestController) Get(context *gin.Context) {
	name := context.Param("name")

	model := c.store.Get(name)

	if model == nil {
		Abort(context, 404, "Manifest not found")
		return
	}

	context.JSON(200, model)
}

func (c *ManifestController) Send(context *gin.Context) {
	name := context.Param("name")

	model := c.store.Get(name)

	if model == nil {
		Abort(context, 404, "Manifest not found")
		return
	}

	var jsonData map[string]interface{}
	if err := context.ShouldBindJSON(&jsonData); err != nil {
		Abort(context, 400, "Invalid JSON")
		return
	}
	value, err := c.services.Component.Validate(model.Components, jsonData)

	if err != nil {
		Abort(context, 500, err.Error())
		return
	}
	context.JSON(200, gin.H{"status": value})

}
