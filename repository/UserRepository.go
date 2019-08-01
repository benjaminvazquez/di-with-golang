package repository

import (
	"fmt"
)

type IUserRepository interface {
	GetUserById(id int) *User
}

type UserRepository struct {
	Configuration *Config
}

func (repository *UserRepository) GetUserById(id int) *User {
	fmt.Println(repository.Configuration.DatabasePath + " from repository")
	user := new(User)
	user.Name = "Benjamin Vazquez"
	user.UserName = "benjaminvazquez"
	user.Id = id
	return user
}

func NewUserRepository(config *Config) IUserRepository {
	rep := new(UserRepository)
	rep.Configuration = config
	return rep
}
