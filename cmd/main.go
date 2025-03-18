package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"cmd/main.go/internal/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/contacts", handlers.GetContacts)
	router.GET("/groups", handlers.GetGroups)

	router.POST("/contacts", handlers.PostContact)
	router.POST("/groups", handlers.PostGroup)

	router.PUT("/contacts/:id/:key/:value", handlers.UpdateContact)
	router.PUT("/groups/:id/:key/:value", handlers.UpdateGroup)

	router.DELETE("/contacts/:id", handlers.DeleteContact)
	router.DELETE("/groups/:id", handlers.DeleteGroup)

	port := os.Getenv("API_PORT")
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}