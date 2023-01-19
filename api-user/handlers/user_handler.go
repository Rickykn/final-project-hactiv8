package handlers

import (
	"api-user/controllers"
	"api-user/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Createuser(c *gin.Context) {

	var userInput *models.Person

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	result := controllers.CreateUser(userInput)

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Success Create User",
		"status code": http.StatusCreated,
		"data":        result,
	})
}
