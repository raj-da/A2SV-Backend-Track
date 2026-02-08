package data

import (
	"context"
	"task-manager/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var TaskCollection *mongo.Collection

func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	var tasks []models.Task
	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return []models.Task{}, err
	}
	defer cursor.Close(context.TODO())
	cursor.All(context.TODO(), &tasks)
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	objID, _ := bson.ObjectIDFromHex(id)
	var task models.Task
	
	err := TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	return task, err
}

func CreateTask(task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := TaskCollection.InsertOne(ctx, task)
	if err == nil {
		task.ID = res.InsertedID.(bson.ObjectID)
	}
	return task, err
}

func UpdateTask(id string, task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	objID, _ := bson.ObjectIDFromHex(id)

	_, err := TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": task,
	})
	return err
}

func DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() 

	objID, _ := bson.ObjectIDFromHex(id)
	
	_, err := TaskCollection.DeleteOne(ctx, bson.M{"_id":objID})
	return err
}