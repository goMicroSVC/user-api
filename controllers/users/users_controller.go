package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goMicroSVC/user-api/models/users"
)

// CreateUser to create a user
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
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
