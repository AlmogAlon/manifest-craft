package controllers

import (
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

type ComponentController struct {
	store    storage.Storage
	services services.Services
}

func NewComponentController(store storage.Storage, services services.Services) *ComponentController {
	return &ComponentController{
		store:    store,
		services: services,
	}
}

func (c *ComponentController) GetValues(context *gin.Context) {
	source := context.Param("source")

	component := c.store.GetComponent(source)

	if component == nil {
		Abort(context, 404, "Component not found")
		return
	}
	values, err := c.services.Component.GetComponentOptions(component)

	if err != nil {
		Abort(context, 500, "an error occurred")
		return
	}
	context.JSON(200, gin.H{"values": values})
}
