package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env tidak ditemukan, menggunakan environment variable")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := gin.Default()

	server.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Go",
		})
	})

	log.Println("Server running on port:", port)

	server.Run(":" + port)
}
