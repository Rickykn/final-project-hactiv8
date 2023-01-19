package handlers

import (
	"api-account/controllers"
	"api-account/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {

	var userInput *models.Account

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	fmt.Println(userInput)
	fmt.Println("Hello")

	result := controllers.CreateAccount(userInput)

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Success Create Account",
		"status code": http.StatusCreated,
		"data":        result,
	})
}
