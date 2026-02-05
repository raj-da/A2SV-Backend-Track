package main

import (
	"log"
	"task-manager/data"
	"task-manager/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	data.ConnectDB()

	// start router
	router := router.SetupRouter()
	router.Run()
}
