package controller

import (
	"MassX/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	res := service.Register(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "failure",
		})
	}
}

func Login(ctx *gin.Context) {
	res := service.Login(ctx)
	if res {
		ctx.SetCookie("username", ctx.PostForm("username"), 240, "/", "127.0.0.1", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "欢迎回来" + ctx.PostForm("username"),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "用户名或密码错误！",
		})
	}
}

func SendMessage(ctx *gin.Context) {
	res := service.SendMessage(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "failure",
		})
	}
}

func ViewMessage(ctx *gin.Context) {
	messages := service.ViewMessage(ctx)
	fmt.Println(messages)
	if messages != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     10000,
			"messages": messages,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": ctx.PostForm("username") + "未留言过",
		})
	}
}