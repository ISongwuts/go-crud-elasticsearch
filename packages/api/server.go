package api

import (
	"github.com/ISongwuts/go-crud-elasticsearch/packages/api/routes"
	"github.com/gin-gonic/gin"
)

func Setup(gin *gin.Engine){
	publicRouter := gin.Group("/api")
	routes.BookRouter(publicRouter)
}