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
	store storage.Storage
}

func NewManifestController(s storage.Storage) *ManifestController {
	return &ManifestController{
		store: s,
	}
}

func (h *ManifestController) Get(context *gin.Context) {
	name := context.Param("name")

	manifest := h.store.Get(name)

	if manifest == nil {
		log.Error("Could not get Manifest ", name)
		context.AbortWithStatusJSON(404, gin.H{"error": "Manifest not found"})
		return
	}

	context.JSON(200, manifest)
}

func (h *ManifestController) Send(context *gin.Context) {
	name := context.Param("name")

	manifest := h.store.Get(name)

	if manifest == nil {
		log.Error("Could not get Manifest ", name)
		context.AbortWithStatusJSON(404, gin.H{"error": "Manifest not found"})
		return
	}

	bodyAsByteArray, _ := ioutil.ReadAll(context.Request.Body)
	jsonMap := make(map[string]string)
	json.Unmarshal(bodyAsByteArray, &jsonMap)

	manifestService := &services.ManifestService{}

	context.JSON(200, gin.H{"status": manifestService.Validate(manifest, jsonMap)})

}
