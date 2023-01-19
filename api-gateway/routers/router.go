package routers

import (
	"api-gateway/handlers"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	engine := gin.New()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.POST("/register", handlers.Register)
	return engine
}
