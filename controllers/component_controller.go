package controllers

import (
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

type ComponentController struct {
	store storage.Storage
}

func NewComponentController(s storage.Storage) *ComponentController {
	return &ComponentController{
		store: s,
	}
}

func (h *ComponentController) GetValues(context *gin.Context) {
	source := context.Param("source")

	component := h.store.GetComponent(source)

	if component == nil {
		context.AbortWithStatusJSON(404, gin.H{"error": "Component not found"})
		return
	}

	componentService := &services.ComponentService{}

	context.JSON(200, gin.H{"values": componentService.GetOptions(component)})
}
