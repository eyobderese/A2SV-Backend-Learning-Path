package router

import (
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/controller"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTaskById)
	router.POST("/tasks", controller.CreateTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)
	router.POST("/register", controller.SignUp)
	router.POST("/login", controller.LogIn)
	router.PUT("/promote/:id", middleware.AdminMiddleware(), controller.Promot)

	return router

}
