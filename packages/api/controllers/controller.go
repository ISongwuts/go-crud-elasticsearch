package controllers

import (
	"github.com/gin-gonic/gin"
	book "github.com/ISongwuts/go-crud-elasticsearch/internal/books/usecase"
	user "github.com/ISongwuts/go-crud-elasticsearch/internal/users/usecase"
)

type (
	IController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetOnce(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	Usecase interface{
		*book.BookUsecase | *user.UserUsecase
	}
)

func NewController[T Usecase](usecase T) IController {
	switch u := any(usecase).(type) {
	case *book.BookUsecase:
		return &BookController{ BookUsecase: u }
	case *user.UserUsecase:
		return &UserController{ UserUsecase: u }
	default:
		return nil
	}
}