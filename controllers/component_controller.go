package controllers

import (
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

type ComponentHandler struct {
	store storage.Storage
}

func NewComponentHandler(s storage.Storage) *ComponentHandler {
	return &ComponentHandler{
		store: s,
	}
}

func (h *ComponentHandler) GetValues(context *gin.Context) {
	source := context.Param("source")

	component := h.store.GetComponent(source)

	if component == nil {
		context.AbortWithStatusJSON(404, gin.H{"error": "Component not found"})
		return
	}

	componentService := &services.ComponentService{}

	context.JSON(200, gin.H{"values": componentService.GetOptions(component)})
}
