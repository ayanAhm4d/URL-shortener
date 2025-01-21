package main

import (
	"log"

	"github.com/ayanAhm4d/URL-shortener/api/database"
	"github.com/ayanAhm4d/URL-shortener/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRoutes(router *gin.Engine, db *database.MySQLStore) {
	router.GET("/:url", func(c *gin.Context) {
		routes.ResolveURL(c, db)
	})
	router.POST("/api/v1", func(c *gin.Context) {
		routes.ShortenURL(c, db)
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewMySQLStore()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	router := gin.Default()
	setupRoutes(router, db)

	log.Fatal(router.Run(":3000"))
}
