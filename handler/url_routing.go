package handler

import (
	"github.com/goMicroSVC/user-api/controllers/ping"
	"github.com/goMicroSVC/user-api/controllers/users"
)

func urlMapping() {
	router.GET("/ping", ping.Ping)
	//router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users/create", users.CreateUser)
	router.Run()
}
