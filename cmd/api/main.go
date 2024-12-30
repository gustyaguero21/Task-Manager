package main

import (
	"task-manager-app/cmd/config"
	"task-manager-app/internal/router"
)

func main() {
	router := router.SetupRouter()

	router.Run(config.Port)
}
