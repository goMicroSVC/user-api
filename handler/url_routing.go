package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/goMicroSVC/user-api/controllers/ping"
	"github.com/goMicroSVC/user-api/controllers/users"
)

var (
	router = gin.Default()
)

func urlMapping() {
	router.GET("/ping", ping.Ping)
	//router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users/create", users.CreateUser)
	router.Run()
}
