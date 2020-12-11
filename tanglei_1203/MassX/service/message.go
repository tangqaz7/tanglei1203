package service

import (
	"MassX/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SendMessage(ctx *gin.Context) bool {
	username, err := ctx.Cookie("username")
	if err != nil {
		return false
	}
	fmt.Println(username)
	message := ctx.PostForm("message")
	res := models.SaveMessage(username, message)
	return res
}

func ViewMessage(ctx *gin.Context) string {
	username := ctx.PostForm("username")
	messages := models.ViewMessage(username)
	bytes, _ := json.Marshal(messages)
	return string(bytes)
}