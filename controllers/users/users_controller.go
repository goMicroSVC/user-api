package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser to create a user
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me :)")
}

// GetUser to Get a user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me :)")
}

// SearchUser to find a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me :)")
}
