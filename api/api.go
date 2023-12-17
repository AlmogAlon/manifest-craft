package api

import (
	"manifest-craft/routes"
	"manifest-craft/services"
	"manifest-craft/storage"

	"github.com/gin-gonic/gin"
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
	appServices := services.Get()

	routes.ManifestRoute(router, s.store, appServices)
	routes.ComponentRoute(router, s.store, appServices)

	return router.Run(s.listenPort)
}
