package main

import (
	"elk-practice/router"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"go.elastic.co/ecslogrus"
)

func main() {

	log.SetFormatter(&ecslogrus.Formatter{})
	log.SetLevel(log.TraceLevel)

	logFilePath := "logs/out.log"
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	}

	defer file.Close()

	fmt.Println("Starting Service")
	log.Info("Starting Service!")

	router := router.NewRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}
}
