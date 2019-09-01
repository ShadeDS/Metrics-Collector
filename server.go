package main

import (
	"github.com/gorilla/mux"
	"log"
	"metrics-collector/controller"
	"metrics-collector/dao"
	"metrics-collector/service"
	"net/http"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "main: ", log.Lshortfile)
}

func main() {
	logger.Println("Start application")
	d := dao.New()
	s := service.New(d)
	c := controller.New(s)

	logger.Println("Initialize http router")
	router := controller.CreateApi(mux.NewRouter(), c)
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Panic(err.Error())
	}
}
