package controller

import (
	"net/http"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/data"
	"github.com/gin-gonic/gin"
)

// GetTasks retrieves all tasks from the data source and returns them as JSON.
func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": tasks})
}

// GetTaskById retrieves a task by its ID.
// It takes a gin.Context object as a parameter and uses the ID parameter from the request path to fetch the task from the data package.
// If the task is found, it returns the task data as a JSON response with HTTP status code 200.
// If the task is not found, it returns a JSON response with HTTP status code 404 and a message indicating that the task was not found.
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": task})
}

// CreateTask creates a new task based on the JSON data provided in the request body.
// It binds the JSON data to a Task struct, creates the task using the data, and returns the created task as JSON response.
func CreateTask(c *gin.Context) {
	var task data.Task
	c.BindJSON(&task)

	task, err := data.CreateTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": task})
}

// UpdateTask updates a task with the given ID.
// It receives the task ID from the request parameters and the updated task data from the request body.
// It returns the updated task if successful, or a JSON response with an error message if the task is not found.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task data.Task
	c.BindJSON(&task)

	task, err := data.UpdateTask(task, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask deletes a task with the given ID.
// It takes a gin.Context object as a parameter and retrieves the task ID from the URL parameter.
// It then calls the data.DeleteTask function to delete the task with the given ID.
// If the task is not found, it returns a JSON response with a "Task not found" message and a status code of 404 (Not Found).
// If the task is successfully deleted, it returns a JSON response with the deleted task data and a status code of 200 (OK).
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := data.DeleteTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"data": task})
}
