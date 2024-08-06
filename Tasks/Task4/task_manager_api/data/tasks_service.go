package data

import (
	"time"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/model"
)

type Task = model.Task

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

// GetTasks returns a slice of tasks.
func GetTasks() []Task {

	return tasks

}

// GetTaskById retrieves a task from the tasks slice based on the provided ID.
// If a task with the given ID is found, it is returned. Otherwise, an empty task is returned.
func GetTaskById(id string) Task {

	for _, task := range tasks {
		if task.ID == id {
			return task
		}
	}
	return Task{}
}

// CreateTask creates a new task and adds it to the list of tasks.
// It takes a Task object as a parameter and returns the created task.
func CreateTask(task Task) Task {

	tasks = append(tasks, task)
	return task
}

// UpdateTask updates a task with the given ID in the tasks slice.
// It replaces the existing task with the provided task and returns the updated task.
// If no task with the given ID is found, it returns an empty Task struct.
func UpdateTask(task Task, id string) Task {

	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = task
			return task

		}
	}
	return Task{}

}

// DeleteTask deletes a task with the specified ID from the tasks slice.
// It returns the deleted task if found, otherwise it returns an empty task.
func DeleteTask(id string) Task {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return t
		}
	}
	return Task{}
}

/*

func EditTask(task Task, id string) Task {

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		return Task{}
	}

	if task.ID != "" {
		tasks[index].ID = task.ID
	}
	if task.Title != "" {

		tasks[index].Title = task.Title
	}
	if task.Description != "" {
		tasks[index].Description = task.Description
	}
	if (task.DueDate != time.Time{}) {
		tasks[index].DueDate = task.DueDate
	}
	if task.Status != "" {
		tasks[index].Status = task.Status
	}
	return tasks[index]
}
*/
