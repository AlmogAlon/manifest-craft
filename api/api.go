package api

import (
	"github.com/gin-gonic/gin"
	"manifest-craft/api/middlewares"
	"manifest-craft/routes"
	"manifest-craft/services"
	"manifest-craft/storage"
)

type Server struct {
	listenPort string

	store storage.Storage
}

func NewServer(listenPort string, store storage.Storage) *Server {
	return &Server{listenPort: listenPort, store: store}

}

func (s *Server) Start() error {
	router := gin.Default()

	middlewares.Use(router)

	appServices := services.Get()

	routes.ManifestRoute(router, s.store, appServices)
	routes.ComponentRoute(router, s.store, appServices)

	return router.Run(s.listenPort)
}
