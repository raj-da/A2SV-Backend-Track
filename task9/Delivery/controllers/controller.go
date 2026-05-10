package controllers

import (
	"net/http"
	domain "task-manager/Domain"

	"github.com/gin-gonic/gin"
)

//* --- --- --- --- --- ---//
//*      Task Controller   //
//* --- --- --- --- --- ---//
type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUsecase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := tc.TaskUsecase.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task Deleted"})
}

func (tc *TaskController) Create(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})		
		return
	}

	if err := tc.TaskUsecase.Create(c, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (tc *TaskController) Update(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TaskUsecase.Update(c, id, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}


//* --- --- --- --- --- ---//
//*     User Controller    //
//* --- --- --- --- --- ---//

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if err := uc.UserUsecase.Register(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration Successful"})
}

func (uc *UserController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := uc.UserUsecase.Login(c, user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// set the refresh token in HttpOnly Cookie
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/refresh", "", true, true)

	c.JSON(http.StatusOK, gin.H{"token": accessToken})
}

func (uc *UserController) Promote(c *gin.Context) {
	username := c.Param("u")
	if err := uc.UserUsecase.PromoteUser(c, username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token missing"})
		return
	}

	newAccessToken, err := uc.UserUsecase.Refresh(c, cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}