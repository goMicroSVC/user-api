package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping func to health check
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong\n")
}
