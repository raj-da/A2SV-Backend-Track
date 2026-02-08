package data

import (
	"context"
	"task-manager/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection

func RegisterUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	count, _ := UserCollection.CountDocuments(ctx, bson.M{})
	// 1. Check if first user
	if count == 0 {
		user.Role = "Admin"
	} else {
		user.Role = "user"
	}

	// 2. Hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)

	// 3. Store user 
	_, err := UserCollection.InsertOne(context.TODO(), user)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var user models.User
	
	err := UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

func PromotUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := UserCollection.UpdateOne(
		ctx, 
		bson.M{"username": username}, 
		bson.M{"$set": bson.M{"role": "Admin"}},
	)
	return err
}