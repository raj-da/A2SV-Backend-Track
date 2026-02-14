package main

import (
	"log"
	"os"
	"task-manager/Delivery/routers"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	db := client.Database("task_manager_3")
	router := routers.SetUp(db)

	router.Run()
}