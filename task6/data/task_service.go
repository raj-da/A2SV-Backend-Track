package data

import (
	"context"
	"errors"
	"log"
	"os"
	"task-manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collection *mongo.Collection

// ConnectDB initialize the MongoDB connection 
func ConnectDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions) // TODO: add context after fixing import issue
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	collection = client.Database("task_db").Collection("tasks")
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var task models.Task
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	return task, err
}

func CreateTask(task models.Task) (models.Task, error) {
	res, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.Task{}, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID)
	return task, nil
}

func UpdateTask(id string, updatedTask models.Task) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{
		"$set": bson.M{
			"title": 		updatedTask.Title,
			"description":  updatedTask.Description,
			"due_date": 	updatedTask.DueDate,
			"status": 		updatedTask.Status,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if result.MatchedCount == 0 {
		return errors.New("no task found to update")
	}
	return err
}

func DeleteTask(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if result.DeletedCount == 0 {
		return errors.New("no task found to delete")
	}
	return err
}