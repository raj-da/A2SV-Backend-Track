package main

import (
	"task4/concurrency"
	"task4/controllers"
	"task4/models"
	"task4/services"
)

func main() {
	// Initialize the service
	lib := services.NewLibrary()

	// Dummy members
	lib.Members[1] = models.Member{ID: 1, Name: "John Doe"}

	// Dummy Book
	lib.AddBook(models.Book{ID: 101, Title: "Go Concurrency", Author: "Rob Pike", Status: "Available"})

	// Create the channel
	reserveChan := make(chan concurrency.ReservationRequest, 10)

	// Start the background worker
	go concurrency.ProcessReservations(lib, reserveChan)

	// Initialize the controller
	controller := controllers.NewLibraryController(lib)

	// Run the app
	controller.Run()
}
