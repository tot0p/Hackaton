package controller

import "github.com/gin-gonic/gin"

func LoginAPIController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
