package main

import (
	"ai-chatbot/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Загружаем переменные из .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл")
	}

	router := gin.Default()

	router.POST("/chat", handlers.ChatHandler)
	router.Static("/", "./web") // если используем веб-интерфейс

	router.Run(":8080")
}
