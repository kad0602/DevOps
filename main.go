package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Створюємо базовий роутер Gin
	r := gin.Default()

	// Налаштовуємо простий ендпоінт
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello World from Go Gin API!",
		})
	})

	// Запускаємо сервер ТА перевіряємо наявність помилок
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Помилка при запуску сервера: %v", err)
	}
}
