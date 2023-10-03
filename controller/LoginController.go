package controller

import "github.com/gin-gonic/gin"

func LoginController(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{
		"title": "Login",
	})
}
