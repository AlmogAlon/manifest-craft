package controllers

import (
	"encoding/json"
	"io/ioutil"
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

	bodyAsByteArray, _ := ioutil.ReadAll(context.Request.Body)
	jsonMap := make(map[string]string)
	err := json.Unmarshal(bodyAsByteArray, &jsonMap)
	if err != nil {
		log.Error("Could not get request info")
		context.AbortWithStatusJSON(404, gin.H{"error": "bad body given"})
		return
	}

	context.JSON(200, gin.H{"status": c.services.Component.Validate(model.Components, jsonMap)})

}
