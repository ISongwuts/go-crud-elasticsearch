package usecase

import (
	"github.com/ISongwuts/go-crud-elasticsearch/config"
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

func (bu *BookUsecase) Create(users *models.Book) (string, error) {
	return "", nil
}

func (bu *BookUsecase) GetBooks() ([]models.Book, error) {
	client, err := config.NewConfig().ElasticSearchConfig()

	if err != nil {
		return nil, err
	}

	books, err := bu.BookRepository.GetBooks(client)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (bu *BookUsecase) GetByID(id string) (models.Book, error) {
	user := models.Book{}
	return user, nil
}

func (bu *BookUsecase) Update(id string, modify map[string]string) (string, error){
	return "", nil
}

func (bu *BookUsecase) Delete(id string) (string, error) {
	return "", nil
}