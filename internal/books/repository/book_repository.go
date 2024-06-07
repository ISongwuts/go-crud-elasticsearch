package repository

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/models"
	"github.com/gin-gonic/gin"
)

type (
	IBookRepository interface {
		Create(user *models.Book)
		GetBooks(ctx *gin.Context) 
		GetByID(id string) 
		Update(id string, modify map[string]string) 
		Delete(id string) 
	}

	BookRepository struct {}
)

func NewBookRepository() IBookRepository {
	return &BookRepository{}
}

func (u *BookRepository) Create(users *models.Book) {

}

func (u *BookRepository) GetBooks(ctx *gin.Context) {
	//users := make([]models.Book, 0)

}

func (u *BookRepository) GetByID(id string) {
	//user := models.Book{}

}

func (u *BookRepository) Update(id string, modify map[string]string) {

}

func (u *BookRepository) Delete(id string) {
}