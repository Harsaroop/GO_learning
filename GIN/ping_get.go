package GIN

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"Updog":    "Whats up dog",
		"Not much": "Whats up with you?",
	})
}
