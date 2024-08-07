package main

import (
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/data"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/router"
)

// main is the entry point of the program.
// It sets up the router and starts the server on localhost:8080.
func main() {
	data.Init()

	router := router.SetupRouter()
	router.Run("localhost:8080")
}
