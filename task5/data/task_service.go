package data

import (
	"errors"
	"task-manager/models"
	"time"
)

// Dummy task data
var tasks = []models.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func CreateTask(task models.Task) {
	tasks = append(tasks, task)
}

func UpdateTask(id string, updatedTask models.Task) error {
	for ind, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[ind] = updatedTask
			return  nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id string) error {
	for ind, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:ind], tasks[ind+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}