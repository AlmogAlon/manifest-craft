package main

import (
	"fmt"
	"manifest-craft/api"
	"manifest-craft/database"
	"manifest-craft/storage"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()

	listenPort := os.Getenv("LISTEN_PORT")
	fmt.Println("LISTEN_PORT:", listenPort)

	// store := storage.NewMemoryStorage()

	store := storage.NewPostgressStorage()
	server := api.NewServer(":"+listenPort, store)

	fmt.Println("Starting server on", listenPort)
	server.Start()
}
