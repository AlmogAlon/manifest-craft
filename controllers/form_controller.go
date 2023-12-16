package controllers

import (
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
)

type FormsHandler struct {
	store storage.Storage
}

func NewFormsHandler(s storage.Storage) *FormsHandler {
	return &FormsHandler{
		store: s,
	}
}

func (h *FormsHandler) Get(context *gin.Context) {
	name := context.Param("name")

	manifest := h.store.Get(name)

	context.JSON(200, manifest)
}
