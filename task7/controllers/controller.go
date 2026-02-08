package controllers

import (
	"net/http"
	"os"
	"task-manager/data"
	"task-manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// User controller
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := data.RegisterUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}
	c.JSON(http.StatusCreated, "User registered")
}

func Login(c *gin.Context) {
	var loginVals models.User
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := data.GetUserByUsername(loginVals.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password)) != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := generateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Promote(c *gin.Context) {
	username := c.Param("username")
	if err := data.PromotUser(username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}

// Helper function
func generateToken(user *models.User) (string, error) {
	expirationTIme := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		Username: user.Username,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTIme),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "before abrham I am",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	return token.SignedString(jwtSecret)
}

// Task Controller
func GetTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	task, err := data.GetTaskByID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := data.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.UpdateTask(c.Param("id"), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}