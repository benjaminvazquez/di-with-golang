package main

import (
	"di-with-golang/controllers"
	"di-with-golang/handlers"
	"di-with-golang/repository"
	"fmt"

	"go.uber.org/dig"
)

func NewConfig() *repository.Config {
	return &repository.Config{
		DatabasePath: "./example.db",
	}
}

func main() {
	container := dig.New()
	container.Provide(NewConfig)
	container.Provide(repository.NewUserRepository)
	container.Provide(controllers.NewUserController)
	container.Provide(handlers.NewUserHandler)

	err := container.Invoke(func(handler *handlers.UserHandler) {
		user := handler.GetUserById(987)
		fmt.Println(user)
	})

	if err != nil {
		panic(err)
	}

}
