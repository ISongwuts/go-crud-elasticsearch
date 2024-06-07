package repository

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/models"
)

type (
	
	IUserRepository interface {
		Create(user *models.User) (string, error)
		GetUsers() ([]models.User, error)
		GetByID(id string) (models.User, error)
		Update(id string, modify map[string]string) (string, error)
		Delete(id string) (string, error)
	}

	UserRepository struct {}
)

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Create(users *models.User) (string, error) {
	return "", nil
}

func (u *UserRepository) GetUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	return users, nil
}

func (u *UserRepository) GetByID(id string) (models.User, error) {
	user := models.User{}
	return user, nil
}

func (u *UserRepository) Update(id string, modify map[string]string) (string, error){
	return "", nil
}

func (u *UserRepository) Delete(id string) (string, error) {
	return "", nil
}