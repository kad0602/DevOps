package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// Тест перевіряє, чи коректно відповідає базовий ендпоінт
func TestHelloWorldRoute(t *testing.T) {
	// Переводимо Gin у тестовий режим, щоб він не спамив у логи під час тестів
	gin.SetMode(gin.TestMode)

	// Створюємо такий самий роутер, як і в main.go
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Hello World from Go Gin API!",
		})
	})

	// Створюємо фейковий HTTP-запит
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	// Виконуємо запит
	r.ServeHTTP(w, req)

	// Перевіряємо, чи статус-код дорівнює 200 (OK)
	if w.Code != http.StatusOK {
		t.Fatalf("Очікувався статус 200, але отримано %d", w.Code)
	}

	// Перевіряємо, чи містить відповідь потрібне слово
	if !strings.Contains(w.Body.String(), "success") {
		t.Fatalf("Тіло відповіді не містить 'success'. Отримано: %s", w.Body.String())
	}
}
