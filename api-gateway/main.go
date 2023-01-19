package main

import (
	"api-gateway/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)

	server := http.Server{
		Addr:    ":8085",
		Handler: routers.Server(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Server Listen Error : %s ", err.Error())
		}
	}()

	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("Server is Shutdown")
}
