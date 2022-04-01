package service

import (
	"log"
	"my_go/pkg/models"
	"my_go/pkg/repository"
)

//Business logic implementation

type userService struct {
	userRepository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	FindAll() []models.User
	ValidateUser(id string) bool
}

func (u userService) FindAll() []models.User {
	//TODO implement me
	var users []models.User
	users, _ = u.userRepository.FindAll()
	return users
}

func InitUserInterface(i repository.UserRepositoryInterface) UserServiceInterface {
	return userService{userRepository: i}
}

func (u userService) ValidateUser(id string) bool {
	//TODO implement me
	user, _ := u.userRepository.FindById(id)

	if &user == nil {
		return false
	}

	log.Print("result-> " + user.Username)
	return true

}
