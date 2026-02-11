package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

//* --- --- --- ---//
//*    Entites     //
//* --- --- --- ---//
// Task entity
type Task struct {
	ID 			bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Title 		string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	DueDate 	time.Time     `json:"due_date" bson:"due_date"`
	Status 		string        `json:"status" bson:"status"`
}

// User entity
type User struct {
	ID 		 bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username" binding:"required"`
	Password string        `json:"password" bson:"password" binding:"required"`
	Role 	 string        `json:"role" bson:"role"`
}

//* --- --- --- --- --- ---//
//*   Repository Interface //
//* --- --- --- --- --- ---//
// Task repository interface
type TaskRepository interface {
	Create(ctx context.Context, task Task) error
	GetByID(ctx context.Context, id string) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, id string, task Task) error
	Delete(ctx context.Context, id string) error
}

// User repository interface
type UserRepository interface {
	Create(ctx context.Context, user User) error
	GetByUsername(ctx context.Context, username string) (User, error)
	UpdateRole(ctx context.Context, username string, role string) error
	// Count(ctx context.Context) (int64, error)
}

//* --- --- --- --- --- ---//
//*     Usecase Interface  //
//* --- --- --- --- --- ---//
// Task usecase interface
type TaskUsecase interface {
	Create(ctx context.Context, task Task) error
	GetByID(ctx context.Context, id string) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, id string, task Task) error
	Delete(ctx context.Context, id string) error
}

// User usecase interface
type UserUsecase interface {
	Register(ctx context.Context, user User) error
	Login(ctx context.Context, username, password string) (string, error)
	PromoteUser(ctx context.Context, username string) error
}