package controller

import (
	"net/http"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/data"
	"github.com/gin-gonic/gin"
)

type User = data.User

func SignUp(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Invalid Content"})
		return
	}

	newuser, errr := data.CreateUser(user)

	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newuser})
}

func LogIn(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Invalid Content"})
		return
	}

	token, errr := data.LoginUser(user)

	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	// create "authorithetion key with token cocke"

	c.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"Authorization": token})

}

func Promot(c *gin.Context) {
	userId := c.Param("id")

	newuser, errr := data.PromoteUser(userId)

	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newuser})
}
