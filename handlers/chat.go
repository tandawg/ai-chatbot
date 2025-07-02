package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

func ChatHandler(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Message) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Мок-ответ, который можно заменить на что угодно
	response := generateMockResponse(req.Message)

	logInteraction(req.Message, response)
	c.JSON(http.StatusOK, ChatResponse{Response: response})
}

func generateMockResponse(message string) string {
	lower := strings.ToLower(message)

	switch {
	case strings.Contains(lower, "sun") && (strings.Contains(lower, "age") || strings.Contains(lower, "old") || strings.Contains(lower, "years")):
		return "The Sun is approximately 4.6 billion years old."
	case strings.Contains(lower, "sun") && strings.Contains(lower, "distance"):
		return "The distance from Earth to Sun is about 149.6 million kilometers."
	case strings.Contains(lower, "earth") && (strings.Contains(lower, "age") || strings.Contains(lower, "old") || strings.Contains(lower, "years")):
		return "The Earth is about 4.54 billion years old."
	case strings.Contains(lower, "earth") && strings.Contains(lower, "shape"):
		return "The Earth is an oblate spheroid."
	case strings.Contains(lower, "hello"):
		return "Hi! How can I help you today?"
	default:
		return "I'm a mock AI! Ask me about the Sun or the Earth."
	}
}

func logInteraction(question, answer string) {
	f, err := os.OpenFile("chatlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	logLine := "Q: " + question + "\nA: " + answer + "\n---\n"
	if _, err := f.WriteString(logLine); err != nil {
		log.Println("Error writing to log file:", err)
	}
}
