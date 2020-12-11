package cmd

import (
	"MassX/controller"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	router := gin.Default()
	router.POST("/register",controller.Register)
	router.GET("/login",controller.Login)
	router.POST("/postMessage",controller.SendMessage)
	router.GET("/viewMessage",controller.ViewMessage)
	router.Run(":8080")
}
