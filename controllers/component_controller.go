package controllers

import (
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

type ComponentController struct {
	store    storage.Storage
	services *services.Services
}

func NewComponentController(store storage.Storage, services *services.Services) *ComponentController {
	return &ComponentController{
		store:    store,
		services: services,
	}
}

func (c *ComponentController) GetValues(context *gin.Context) {
	source := context.Param("source")

	component := c.store.GetComponent(source)

	if component == nil {
		context.AbortWithStatusJSON(404, gin.H{"error": "Component not found"})
		return
	}

	context.JSON(200, gin.H{"values": c.services.Component.GetComponentOptions(component)})
}
