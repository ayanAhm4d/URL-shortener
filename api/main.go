package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ayanAhm4d/URL-shortener/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// setup two routes, one for shortening the url
// the other for resolving the url
// for example if the short is `4fg`, the user
// must navigate to `localhost:3000/4fg` to redirect to
// original URL. The domain can be changes in .env file
func setupRoutes(router *gin.Engine) {
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/v1", routes.ShortenURL)
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	setupRoutes(router)

	log.Fatal(router.Run(os.Getenv("APP_PORT")))
}
