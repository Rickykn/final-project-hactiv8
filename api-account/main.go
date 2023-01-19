package main

import (
	"api-account/database"
	"api-account/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	errConnect := database.Connect()

	if errConnect != nil {
		panic(errConnect)
	}
	gin.SetMode(gin.DebugMode)

	server := http.Server{
		Addr:    ":8084",
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
