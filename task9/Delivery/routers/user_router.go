package routers

import (
	"task-manager/Delivery/controllers"
	infrastructure "task-manager/Infrastructure"
	repository "task-manager/Repository"
	usecases "task-manager/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRouter(group *gin.RouterGroup, db *mongo.Database, collection []string) {
	// 1. Create the data layer
	userRepo := repository.NewUserRepository(db, collection[0])
	refreshTokenRepo := repository.NewRefreshTokenRepo(db, collection[1])

	// 2. Create infrastructure structs
	jwtService := infrastructure.NewJWTService()
	tokenService := infrastructure.NewTokenService()

	// 3. Create the usecase layer
	useCase := usecases.NewUserUsecase(
		userRepo, 
		refreshTokenRepo, 
		jwtService, 
		tokenService,
	)

	// 4. Create the controller layer
	uc := &controllers.UserController{UserUsecase: useCase}

	// 5. Map the paths
	group.POST("/register", uc.Register)
	group.POST("/login", uc.Login)
	group.PATCH("/promote/:u", uc.Promote).Use(infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware())
	group.POST("/refresh", uc.RefreshToken)
}