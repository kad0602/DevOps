package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// Зчитуємо секрет, який Cloud Run передасть у змінну оточення
		secretKey := os.Getenv("MY_API_KEY")
		if secretKey == "" {
			secretKey = "Секрет локальний або не знайдено!"
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello World from Google Cloud Run!",
			"secret":  secretKey, // Виводимо для демонстрації викладачу
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Помилка при запуску сервера: %v", err)
	}
}
