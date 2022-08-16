package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juhonamnam/gin-practice.git/entity"
	"github.com/juhonamnam/gin-practice.git/service"
)

type ToDoController interface {
	FindAll() []entity.ToDo
	Save(ctx *gin.Context) entity.ToDo
}

type controller struct {
	service service.ToDoService
}

func New(service service.ToDoService) ToDoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.ToDo {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.ToDo {
	var toDo entity.ToDo
	ctx.BindJSON(&toDo)
	c.service.Save(toDo)
	return toDo
}
