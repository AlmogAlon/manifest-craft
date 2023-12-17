package controllers

import (
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

type ManifestHandler struct {
	store storage.Storage
}

func NewManifestHandler(s storage.Storage) *ManifestHandler {
	return &ManifestHandler{
		store: s,
	}
}

func (h *ManifestHandler) Get(context *gin.Context) {
	name := context.Param("name")

	manifest := h.store.Get(name)

	if manifest == nil {
		context.AbortWithStatusJSON(404, gin.H{"error": "Manifest not found"})
		return
	}

	context.JSON(200, manifest)
}
