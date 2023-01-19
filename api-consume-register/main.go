package main

import (
	"api-consume-register/kafka"
	"api-consume-register/routers"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello its consumer service")
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	consumer, err := kafka.ConnConsumer()
	if err != nil {
		fmt.Println("Error Consumer connection ", err.Error())
		os.Exit(1)
	}
	go kafka.Subscriber(consumer)
	gin.SetMode(gin.DebugMode)
	routers.Server().Run(":8083")
}
