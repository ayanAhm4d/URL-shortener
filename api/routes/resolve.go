package routes

import (
	"github.com/ayanAhm4d/URL-shortener/database"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// ResolveURL ...
func ResolveURL(c *gin.Context) {
	// get the short from the url
	url := c.Param("url")
	// query the db to find the original URL, if a match is found
	// increment the redirect counter and redirect to the original URL
	// else return error message
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		c.JSON(404, gin.H{"error": "short not found on database"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": "cannot connect to DB"})
		return
	}
	// increment the counter
	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")
	// redirect to original URL
	c.Redirect(301, value)
}
