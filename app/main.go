package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/ISongwuts/go-crud-elasticsearch/packages/api"
)

func main() {
	server := gin.Default()
	api.Setup(server)
	log.Println("Starting server on port 8000...")
	if err := server.Run(":8000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
