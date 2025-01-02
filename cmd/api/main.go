package main

import (
	"log"
	"task-manager-app/cmd/config"
	"task-manager-app/internal/data"
	"task-manager-app/internal/router"
)

func main() {
	router := router.SetupRouter()

	conn, err := data.InitDatabase()
	if err != nil {
		log.Fatal("error creating database")
	}

	defer conn.Close()

	router.Run(config.Port)
}
