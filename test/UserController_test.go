package test

import (
	"di-with-golang/controllers"
	"di-with-golang/repository"
	"fmt"
	"testing"
)

type FakeRepo struct {
}

func (fake *FakeRepo) GetUserById(id int) *repository.User {
	return &repository.User{Id: id}
}

func TestGetUserById(test *testing.T) {
	repo := &FakeRepo{}
	controller := controllers.NewUserController(repo)
	user := controller.GetUserById(1)
	if user.Id != 1 {
		test.Fatal("The user Id doesn't match")
	}
	fmt.Println(user)
}
