package main

import (
	"manifest-craft/api"
	"manifest-craft/config"
	"manifest-craft/storage"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()

	// config.ConnectToDB()
	config.InitLog()

	listenPort := os.Getenv("LISTEN_PORT")
	log.Info("LISTEN_PORT:", listenPort)

	store := storage.NewMemoryStorage()

	// store := storage.NewPostgressStorage()
	server := api.NewServer(":"+listenPort, store)

	log.Info("Starting server on", listenPort)
	server.Start()
}
