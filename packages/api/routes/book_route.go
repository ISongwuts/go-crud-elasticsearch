package routes

import (
	"github.com/ISongwuts/go-crud-elasticsearch/internal/books/usecase"
	"github.com/ISongwuts/go-crud-elasticsearch/packages/api/controllers"
	"github.com/gin-gonic/gin"
)

func BookRouter(rg *gin.RouterGroup) {
	controller := controllers.NewController(&usecase.BookUsecase{})
	
	//Route paths
	rg.GET("/books", controller.GetAll)
	rg.GET("/book/:id", controller.GetOnce)
	rg.POST("/books", controller.Create)
	rg.PUT("/book/:id", controller.Update)
	rg.DELETE("/book/:id", controller.Delete)
}