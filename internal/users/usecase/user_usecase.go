package usecase

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/models"
	"github.com/ISongwuts/go-crud-elasticsearch/internal/users/repository"
)

type (
	UserUsecase struct {
		repository.UserRepository
	}

	IUserUsecase interface {
		Create(user *models.User) (string, error)
		GetUsers() ([]models.User, error)
		GetByID(id string) (models.User, error)
		Update(id string, modify map[string]string) (string, error)
		Delete(id string) (string, error)
	}
)

func NewUserUsecase() IUserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) Create(users *models.User) (string, error) {
	return "", nil
}

func (u *UserUsecase) GetUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	return users, nil
}

func (u *UserUsecase) GetByID(id string) (models.User, error) {
	user := models.User{}
	return user, nil
}

func (u *UserUsecase) Update(id string, modify map[string]string) (string, error){
	return "", nil
}

func (u *UserUsecase) Delete(id string) (string, error) {
	return "", nil
}