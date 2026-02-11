package repository

import (
	"context"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepository struct {
	db *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, col string) domain.UserRepository {
	return &userRepository{
		db: db,
		collection: col,
	}
}

func (ur *userRepository) Create(ctx context.Context, user domain.User) error {
	_, err := ur.db.Collection(ur.collection).InsertOne(ctx, user)
	return err
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	res := ur.db.Collection(ur.collection).FindOne(ctx, bson.M{"username":username})
	if err := res.Decode(&user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur *userRepository) UpdateRole(ctx context.Context, username, role string) error {
	_, err := ur.db.Collection(ur.collection).UpdateOne(ctx, bson.M{"username": username}, bson.M{
		"$set": bson.M{"role": role},
	})

	return err
}