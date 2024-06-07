package controllers

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/users/usecase"
	"github.com/gin-gonic/gin"
)

type (
	UserController struct{
		UserUsecase *usecase.UserUsecase
	}
)

func (uc *UserController) Create(ctx *gin.Context) {

}

func (uc *UserController) GetAll(ctx *gin.Context) {

}

func (uc *UserController) GetOnce(ctx *gin.Context) {
	
}

func (uc *UserController) Update(ctx *gin.Context) {
	
}

func (uc *UserController) Delete(ctx *gin.Context) {
	
}