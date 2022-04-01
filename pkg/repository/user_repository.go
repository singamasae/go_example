package repository

import (
	"gorm.io/gorm"
	"my_go/pkg/models"
)

//TODO: SQL Query implementation

type userRepository struct {
	Db *gorm.DB
}

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindById(id string) (models.User, error)
	FindAll() ([]models.User, error)
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return userRepository{Db: db}
}

func (u userRepository) Save(user models.User) (models.User, error) {
	//TODO implement me
	err := u.Db.Create(&user).Error
	return user, err
}

func (u userRepository) FindById(id string) (models.User, error) {
	//TODO implement me
	var user = models.User{}
	err := u.Db.Where("id=?", id).Find(&user).Error
	return user, err
}

func (u userRepository) FindAll() (users []models.User, err error) {
	//TODO implement me
	err = u.Db.Find(&users).Error
	return users, err
}
