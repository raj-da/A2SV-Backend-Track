package main

import (
	"task3/controllers"
	"task3/models"
	"task3/services"
)

func main() {
	// Initialize the service
	lib := services.NewLibrary()

	// Dummy members
	lib.Members[1] = models.Member{ID: 1, Name: "John Doe"}

	// Initialize the controller
	controller := controllers.NewLibraryController(lib)

	// Run the app
	controller.Run()
}
