package controllers

import (
	"github.com/gin-gonic/gin"
	book "github.com/ISongwuts/go-crud-elasticsearch/internal/books/usecase"
	user "github.com/ISongwuts/go-crud-elasticsearch/internal/users/usecase"
)

type (
	IController interface {
		GetAll(ctx *gin.Context)
		GetOnce(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	Usecase interface{
		book.BookUsecase | user.UserUsecase
	}
)

func NewController[T Usecase](repo T) IController {
	switch any(repo).(type) {
	case book.BookUsecase:
		return &BookController{}
	case user.UserUsecase:
		return &UserController{}
	default:
		return nil
	}
}