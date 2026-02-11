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

	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (tc *TaskController) Update(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc.TaskUsecase.Update(c, id, task)
}


//* --- --- --- --- --- ---//
//*     User Controller    //
//* --- --- --- --- --- ---//