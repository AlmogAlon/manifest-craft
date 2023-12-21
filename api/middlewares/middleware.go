package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// a custom middleware to handle exceptions globally.
func recoveryMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(fmt.Sprintf("%v", err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}()

	c.Next()
}

func Use(router *gin.Engine) {
	middlewares := []gin.HandlerFunc{
		recoveryMiddleware,
	}

	router.Use(middlewares...)
}
