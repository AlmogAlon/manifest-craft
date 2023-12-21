package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Abort(c *gin.Context, code int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
