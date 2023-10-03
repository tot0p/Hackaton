package controller

import "github.com/gin-gonic/gin"

func IndexController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
