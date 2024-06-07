package controllers

import (
	"github.com/gin-gonic/gin"
	book "github.com/ISongwuts/go-crud-elasticsearch/internal/books/repository"
	user "github.com/ISongwuts/go-crud-elasticsearch/internal/users/repository"
)

type (
	IController interface {
		GetAll(ctx *gin.Context)
		GetOnce(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	Repository interface{
		book.BookRepository | user.UserRepository
	}
)

func NewController[T Repository](repo T) IController {
	switch any(repo).(type) {
	case book.BookRepository:
		return &BookController{}
	case user.UserRepository:
		return &UserController{}
	default:
		return nil
	}
}