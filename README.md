# Manifest Craft

library for checking manifests components through a simple interface

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
      go mod download 
      go run cmd/main.go
    ```
  
## Running tests
- execute these commands:
    ```bash
      cd manifest-craft
      go test ./...
    ```
To run the services using docker-compose:
- execute these commands: 

    ```bash
      cd manifest-craft
      docker-compose build
      docker-compose up
    ```

## Adding new Manifest or Component
- You can currently only add in-mem manifests + components in ```storage/memory.go``` file


- Newly created Component type (source) need to be created under ```services/components/providers/``` directory:

  - implement the Provider interface 
  - register it (source name -> provider) to ```providers.InitProviders``` function
