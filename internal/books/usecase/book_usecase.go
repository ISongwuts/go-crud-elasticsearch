package usecase

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/books/repository"
	"github.com/ISongwuts/go-crud-elasticsearch/internal/models"
)

type (
	BookUsecase struct{
		repository.BookRepository
	}

	IBookUsecase interface {
		Create(user *models.Book) (string, error)
		GetBooks() ([]models.Book, error)
		GetByID(id string) (models.Book, error)
		Update(id string, modify map[string]string) (string, error)
		Delete(id string) (string, error)
	}
)

func NewBookUsecase() IBookUsecase {
	return &BookUsecase{}
}

func (u *BookUsecase) Create(users *models.Book) (string, error) {
	return "", nil
}

func (u *BookUsecase) GetBooks() ([]models.Book, error) {
	users := make([]models.Book, 0)
	return users, nil
}

func (u *BookUsecase) GetByID(id string) (models.Book, error) {
	user := models.Book{}
	return user, nil
}

func (u *BookUsecase) Update(id string, modify map[string]string) (string, error){
	return "", nil
}

func (u *BookUsecase) Delete(id string) (string, error) {
	return "", nil
}