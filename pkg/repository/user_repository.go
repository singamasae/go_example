package repository

import (
	"gorm.io/gorm"
	"log"
	"my_go/pkg/models"
)

type userRepository struct {
	Db *gorm.DB
}

type UserRepositoryInterface interface {
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
}

func InitUserRepository(db *gorm.DB) UserRepositoryInterface {
	return userRepository{Db: db}
}

func (u userRepository) FindAll() (users []models.User, err error) {
	//TODO implement me
	err = u.Db.Find(&users).Error
	return users, err
}

func (u userRepository) FindById(id string) (models.User, error) {
	//TODO implement me
	log.Print("Repo" + id)
	var user = models.User{}
	err := u.Db.Where("id=?", id).Find(&user).Error
	return user, err
}
