package routes

import (
	"math"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ayanAhm4d/URL-shortener/api/database"
	"github.com/ayanAhm4d/URL-shortener/api/helpers"
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
	XRateLimitReset float64       `json:"rate_limit_reset"`
}

func ShortenURL(c *gin.Context, db *database.MySQLStore) {
	var body request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Rate limiting
	remaining, resetAt, err := db.CheckRateLimit(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limit check failed"})
		return
	}

	if remaining < 0 {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error":            "rate limit exceeded",
			"rate_limit_reset": time.Until(resetAt).Minutes(),
		})
		return
	}

	// URL validation
	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL"})
		return
	}

	if !helpers.RemoveDomainError(body.URL) {
		c.JSON(http.StatusForbidden, gin.H{"error": "domain not allowed"})
		return
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	var short string
	if body.CustomShort == "" {
		short = uuid.New().String()[:6]
	} else {
		short = body.CustomShort
	}

	// Check existing short URL
	existing, err := db.GetURL(short)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if existing != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "short URL already exists"})
		return
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	if err := db.SaveURL(short, body.URL, body.Expiry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save URL"})
		return
	}

	c.JSON(http.StatusOK, response{
		URL:             body.URL,
		CustomShort:     "localhost:3000/" + short,
		Expiry:          body.Expiry,
		XRateRemaining:  remaining,
		XRateLimitReset: math.Ceil(time.Until(resetAt).Minutes()),
	})
}
