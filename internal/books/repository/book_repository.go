package repository

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/models"
)

type (
	IBookRepository interface {
		Create(user *models.Book) (string, error)
		GetBooks() ([]models.Book, error)
		GetByID(id string) (models.Book, error)
		Update(id string, modify map[string]string) (string, error)
		Delete(id string) (string, error)
	}

	BookRepository struct {}
)

func NewBookRepository() IBookRepository {
	return &BookRepository{}
}

func (u *BookRepository) Create(users *models.Book) (string, error) {
	return "", nil
}

func (u *BookRepository) GetBooks() ([]models.Book, error) {
	users := make([]models.Book, 0)
	return users, nil
}

func (u *BookRepository) GetByID(id string) (models.Book, error) {
	user := models.Book{}
	return user, nil
}

func (u *BookRepository) Update(id string, modify map[string]string) (string, error){
	return "", nil
}

func (u *BookRepository) Delete(id string) (string, error) {
	return "", nil
}