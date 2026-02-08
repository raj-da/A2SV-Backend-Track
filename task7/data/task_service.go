package data

import (
	"context"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var TaskCollection *mongo.Collection

func GetAllTasks() ([]models.Task, error) {
	//TODO: Add proper context
	var tasks []models.Task
	cursor, err := TaskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Task{}, err
	}
	defer cursor.Close(context.TODO())
	cursor.All(context.TODO(), &tasks)
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	objID, _ := bson.ObjectIDFromHex(id)
	var task models.Task
	//TODO: add proper context
	err := TaskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	return task, err
}

func CreateTask(task models.Task) (models.Task, error) {
	//TODO: add proper context
	res, err := TaskCollection.InsertOne(context.TODO(), task)
	if err == nil {
		task.ID = res.InsertedID.(bson.ObjectID)
	}
	return task, err
}

func UpdateTask(id string, task models.Task) error {
	objID, _ := bson.ObjectIDFromHex(id)
	// TODO: add proper context
	_, err := TaskCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{
		"$set": task,
	})
	return err
}

func DeleteTask(id string) error {
	objID, _ := bson.ObjectIDFromHex(id)
	// TODO: add proper context
	_, err := TaskCollection.DeleteOne(context.TODO(), bson.M{"_id":objID})
	return err
}