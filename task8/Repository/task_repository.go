package repository

import (
	"context"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type taskRepository struct {
	database *mongo.Database
	collection string
}

// Factory function to return taskRepository function
func NewTaskRepository(db *mongo.Database, col string) domain.TaskRepository {
	return &taskRepository{
		database: db,
		collection: col,
	}
}

func (tr *taskRepository) Create(ctx context.Context, task domain.Task) error {
	_, err := tr.database.Collection(tr.collection).InsertOne(ctx, task)
	return  err
}

func (tr *taskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := tr.database.Collection(tr.collection).Find(ctx, bson.M{})

	if err != nil {
		return []domain.Task{}, err
	}
	cursor.All(ctx, &tasks)
	return tasks, nil
}

func (tr *taskRepository) GetByID(ctx context.Context, id string) (domain.Task, error) {
	objID, _ := bson.ObjectIDFromHex(id)
	var task domain.Task
	res := tr.database.Collection(tr.collection).FindOne(ctx, bson.M{"_id": objID})
	if err := res.Decode(&task); err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (tr *taskRepository) Update(ctx context.Context, id string, task domain.Task) error {
	objID, _ := bson.ObjectIDFromHex(id)
	_, err := tr.database.Collection(tr.collection).UpdateByID(ctx, objID, bson.M{"$set": task})
	return err
}

func (tr *taskRepository) Delete(ctx context.Context, id string) error {
	objID, _ := bson.ObjectIDFromHex(id)
	_, err := tr.database.Collection(tr.collection).DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}