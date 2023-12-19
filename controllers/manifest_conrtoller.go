package controllers

import (
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ManifestController struct {
	store    storage.Storage
	services *services.Services
}

func NewManifestController(storage storage.Storage, services *services.Services) *ManifestController {
	return &ManifestController{
		store:    storage,
		services: services,
	}
}

func (c *ManifestController) Get(context *gin.Context) {
	name := context.Param("name")

	model := c.store.Get(name)

	if model == nil {
		log.Error("Could not get Manifest ", name)
		context.AbortWithStatusJSON(404, gin.H{"error": "Manifest not found"})
		return
	}

	context.JSON(200, model)
}

func (c *ManifestController) Send(context *gin.Context) {
	name := context.Param("name")

	model := c.store.Get(name)

	if model == nil {
		log.Error("Could not get Manifest ", name)
		context.AbortWithStatusJSON(404, gin.H{"error": "Manifest not found"})
		return
	}

	var jsonData map[string]interface{}
	if err := context.ShouldBindJSON(&jsonData); err != nil {
		log.Error("Could not get request info")
		context.AbortWithStatusJSON(404, gin.H{"error": "bad body given"})
		return
	}
	value, err := c.services.Component.Validate(model.Components, jsonData)

	if err != nil {
		log.Error("Could not get request info")
		context.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"status": value})

}
