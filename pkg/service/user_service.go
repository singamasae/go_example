package service

import (
	"my_go/pkg/models"
	"my_go/pkg/repository"
)

//TODO: Business logic implementation

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) Save(user models.User) (models.User, error) {
	//TODO implement me
	return u.userRepository.Save(user)
}

func (u userService) FindById(id string) models.User {
	//TODO implement me
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return models.User{}
	}
	return user
}

func (u userService) FindAll() []models.User {
	//TODO implement me
	var users []models.User
	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil
	}
	return users
}

type UserService interface {
	Save(user models.User) (models.User, error)
	FindById(id string) models.User
	FindAll() []models.User
}

func InitUserInterface(i repository.UserRepository) UserService {
	return userService{userRepository: i}
}
