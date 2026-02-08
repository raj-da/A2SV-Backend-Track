package data

import (
	"context"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection

func RegisterUser(user models.User) error {
	// TODO: add correct context
	count, _ := UserCollection.CountDocuments(context.TODO(), bson.M{})
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
	var user models.User
	// TODO: add correct context
	err := UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	return user, err
}

func PromotUser(username string) error {
	// TODO: add correct context
	_, err := UserCollection.UpdateOne(
		context.TODO(), 
		bson.M{"username": username}, 
		bson.M{"$set": bson.M{"role": "Admin"}},
	)
	return err
}