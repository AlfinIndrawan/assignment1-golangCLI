package main

import (
	"assignment3/controller"
	"assignment3/repository"
	"assignment3/service"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		repository repository.Repository = repository.NewDataRepository("database/db.json")
		service    service.Service       = service.NewService(repository)
		controller controller.Controller = controller.NewController(service)
	)
	server := gin.Default()
	Repeat := func(f func() error, delay time.Duration) {
		go func() {
			for {
				f()
				time.Sleep(delay * time.Second)
			}
		}()
	}

	go Repeat(service.UpdateData, 15)
	server.GET("/", controller.GetStatus)

	server.Run(":8080")
}
