package controllers

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/books/usecase"
	"github.com/gin-gonic/gin"
)

type (
	BookController struct{
		BookUsecase *usecase.BookUsecase
	}
)

func (bc *BookController) Create(ctx *gin.Context) {
	
}


func (bc *BookController) GetAll(ctx *gin.Context) {
	
}

func (bc *BookController) GetOnce(ctx *gin.Context) {
	
}

func (bc *BookController) Update(ctx *gin.Context) {
	
}

func (bc *BookController) Delete(ctx *gin.Context) {
	
}