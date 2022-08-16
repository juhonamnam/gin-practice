package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juhonamnam/gin-practice.git/controller"
	"github.com/juhonamnam/gin-practice.git/service"
)

var (
	toDoService    service.ToDoService       = service.New()
	toDoController controller.ToDoController = controller.New(toDoService)
)

func main() {
	server := gin.Default()
	server.GET("/retrieve", func(ctx *gin.Context) {
		ctx.JSON(200, toDoController.FindAll())
	})
	server.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(200, toDoController.Save(ctx))
	})
	server.Run(":8080")
}
