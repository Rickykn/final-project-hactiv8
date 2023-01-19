package main

import (
	"register/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", controller.Register)

	err := r.Run(":8081")

	if err != nil {
		panic(err)
	}
}
