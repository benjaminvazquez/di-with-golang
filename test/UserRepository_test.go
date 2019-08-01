package test

import (
	"di-with-golang/repository"
	"testing"
)

func TestGetUserByIdRepo(test *testing.T) {
	config := &repository.Config{DatabasePath: "Fake database path"}
	repository := repository.NewUserRepository(config)
	user := repository.GetUserById(1)
	if user.Id != 1 {
		test.Fatal("The user Id doesn't match")
	}
}
