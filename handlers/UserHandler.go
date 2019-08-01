package handlers

import (
	"di-with-golang/controllers"
)

type UserHandler struct {
	controller controllers.IUserController
}

func (userHandler *UserHandler) GetUserById(id int) *controllers.User {
	user := userHandler.controller.GetUserById(id)
	return user
}

func NewUserHandler(userController controllers.IUserController) *UserHandler {
	return &UserHandler{controller: userController}
}
