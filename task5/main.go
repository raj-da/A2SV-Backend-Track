package main

import "task-manager/router"

func main() {
	router := router.SetupRouter()
	router.Run()
}