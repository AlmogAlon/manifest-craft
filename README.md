# Manifest Craft

library for creating manifests through a simple interface

- gin as the web framework
- air for live reload
- gorm
- postgres


## Requirements
- Go 1.21
- GoLand or VSCode
- Docker (optional)
- docker-compose (optional)


## Running the application
create .env file with the desired application LISTEN_PORT

To run the application locally:
- execute these commands:
    ```bash
      cd manifest-craft
      go run cmd/main.go
    ```

To run the services using docker-compose:
- execute these commands: 

    ```bash
      cd manifest-craft
      docker-compose build
      docker-compose up
    ```
