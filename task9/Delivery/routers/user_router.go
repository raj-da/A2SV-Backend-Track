package routers

import (
	"task-manager/Delivery/controllers"
	infrastructure "task-manager/Infrastructure"
	repository "task-manager/Repository"
	usecases "task-manager/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRouter(group *gin.RouterGroup, db *mongo.Database, collection string) {
	// 1. Create the data layer
	repo := repository.NewUserRepository(db, collection)

	// 2. Create the usecase layer
	useCase := usecases.NewUserUsecase(repo)

	// 3. Create the controller layer
	uc := &controllers.UserController{UserUsecase: useCase}

	// 4. Map the paths
	group.POST("/register", uc.Register)
	group.POST("/login", uc.Login)
	group.PATCH("/promote/:u", uc.Promote).Use(infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware())
}