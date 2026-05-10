package repository

import (
	"context"
	"errors"
	"fmt"
	"task-manager/Domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type refreshTokenRepo struct {
	database *mongo.Database
	collection string
}

func NewRefreshTokenRepo(db *mongo.Database, col string) *refreshTokenRepo {
	return &refreshTokenRepo{
		database: db,
		collection: col,
	}
}

func (r *refreshTokenRepo) Create(ctx context.Context, token Domain.RefreshToken) error {
	_, err := r.database.Collection(r.collection).InsertOne(ctx, token)
	return err
}

func (r *refreshTokenRepo) GetByToken(ctx context.Context, token string) (Domain.RefreshToken, error) {
	var refreshtToken Domain.RefreshToken

	filter := bson.M{"token": token}
	err := r.database.Collection(r.collection).FindOne(ctx, filter).Decode(&refreshtToken)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Domain.RefreshToken{}, fmt.Errorf("refresh token not found")
		}
		return Domain.RefreshToken{}, fmt.Errorf("failed to get refresh token: %w", err)
	}

	return refreshtToken, nil
}

func (r *refreshTokenRepo) DeleteByToken(ctx context.Context, token string) error {
	filter := bson.M{"token": token}
	result, err := r.database.Collection(r.collection).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("refresh token not found")
	}

	return nil
}

func (r *refreshTokenRepo) DeleteByUserID(ctx context.Context, userID string) error {
	objID, _ := bson.ObjectIDFromHex(userID)
	filter := bson.M{"userID": objID}
	_, err := r.database.Collection(r.collection).DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete refresh tokens for user: %w", err)
	}

	return nil
}