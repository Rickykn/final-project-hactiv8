package handlers

import (
	"api-gateway/host"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var personReq *host.Person

	err := c.ShouldBindJSON(&personReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err2 := host.RegisterGateway(personReq)

	if err2 == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message":     "Register On Process",
			"status code": http.StatusCreated,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Register Failed",
		"status code": http.StatusBadGateway,
	})

}
