package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/goMicroSVC/user-api/controllers"
)

var (
	router = gin.Default()
)

// StartApp func for http handler
func StartApp() {
	router.GET("/ping", controllers.Ping)
	router.Run()
}
