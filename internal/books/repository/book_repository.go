package repository

import (
	"github.com/gin-gonic/gin"
)

type (
	IBookRepository interface {
		Create(ctx *gin.Context)
		GetBooks(ctx *gin.Context) 
		GetByID(ctx *gin.Context) 
		Update(ctx *gin.Context) 
		Delete(ctx *gin.Context) 
	}

	BookRepository struct {}
)

func NewBookRepository() IBookRepository {
	return &BookRepository{}
}

func (u *BookRepository) Create(ctx *gin.Context) {

}

func (u *BookRepository) GetBooks(ctx *gin.Context) {
	//users := make([]models.Book, 0)

}

func (u *BookRepository) GetByID(ctx *gin.Context) {
	//user := models.Book{}

}

func (u *BookRepository) Update(ctx *gin.Context) {

}

func (u *BookRepository) Delete(ctx *gin.Context) {
}