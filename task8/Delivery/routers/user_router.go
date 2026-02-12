package routers

import (
	"task-manager/Delivery/controllers"
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
	group.POST("/user", uc.Register)
	group.POST("/user", uc.Login)
	group.PATCH("/user/:u", uc.Promote)
}