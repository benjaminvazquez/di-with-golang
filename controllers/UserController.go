package controllers

import (
	"di-with-golang/repository"
)

type IUserController interface {
	GetUserById(id int) *User
}

type UserController struct {
	Repository repository.IUserRepository
}

func (userController *UserController) GetUserById(id int) *User {
	user := userController.Repository.GetUserById(id)
	return mapUser(user)
}

func mapUser(model *repository.User) *User {
	newUser := new(User)
	newUser.Name = model.Name
	newUser.UserName = model.UserName
	newUser.Id = model.Id
	return newUser
}

func NewUserController(repository repository.IUserRepository) IUserController {
	controller := new(UserController)
	controller.Repository = repository
	return controller
}
