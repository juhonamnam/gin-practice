package service

import "github.com/juhonamnam/gin-practice.git/entity"

type ToDoService interface {
	Save(entity.ToDo) entity.ToDo
	FindAll() []entity.ToDo
}

type toDoService struct {
	todolist []entity.ToDo
}

func New() ToDoService {
	return &toDoService{
		todolist: []entity.ToDo{},
	}
}

func (service *toDoService) Save(toDoItem entity.ToDo) entity.ToDo {
	service.todolist = append(service.todolist, toDoItem)
	return toDoItem
}

func (service *toDoService) FindAll() []entity.ToDo {
	return service.todolist
}
