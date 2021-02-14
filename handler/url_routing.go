package handler

import "github.com/goMicroSVC/user-api/controllers"

func urlMapping() {
	router.GET("/ping", controllers.Ping)
	router.Run()
}
