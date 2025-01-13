package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ayanAhm4d/URL-shortener/database"
	"github.com/ayanAhm4d/URL-shortener/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

// ShortenURL ...
func ShortenURL(c *gin.Context)  {
	// check for the incoming request body
	body := new(request)
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "cannot parse JSON"})
		return
	}

	// implement rate limiting
	// everytime a user queries, check if the IP is already in database,
	// if yes, decrement the calls remaining by one, else add the IP to database
	// with expiry of `30mins`. So in this case the user will be able to send 10
	// requests every 30 minutes
	r2 := database.CreateClient(1)
	defer r2.Close()
	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
			c.JSON(503, gin.H{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Minute,
			})
			return
		}
	}

	// check if the input is an actual URL
	if !govalidator.IsURL(body.URL) {
		c.JSON(400, gin.H{"error": "Invalid URL"})
		return
	}

	// check for the domain error
	// users may abuse the shortener by shorting the domain `localhost:3000` itself
	// leading to a inifite loop, so don't accept the domain for shortening
	if !helpers.RemoveDomainError(body.URL) {
		c.JSON(503, gin.H{"error": "haha... nice try"})
		return
	}

	// enforce https
	// all url will be converted to https before storing in database
	body.URL = helpers.EnforceHTTP(body.URL)

	// check if the user has provided any custom dhort urls
	// if yes, proceed,
	// else, create a new short using the first 6 digits of uuid
	// haven't performed any collision checks on this
	// you can create one for your own
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	// check if the user provided short is already in use
	if val != "" {
		c.JSON(403, gin.H{"error": "URL short already in use"})
		return
	}

	if body.Expiry == 0 {
		body.Expiry = 24 // default expiry of 24 hours
	}
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to connect to server"})
		return
	}
	// respond with the url, short, expiry in hours, calls remaining and time to reset
	resp := response{
		URL:             body.URL,
		CustomShort:     os.Getenv("DOMAIN") + "/" + id,
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}
	r2.Decr(database.Ctx, c.ClientIP())
	val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
	resp.XRateLimitReset = ttl / time.Minute

	c.JSON(200, resp)
}
