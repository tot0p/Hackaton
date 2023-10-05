package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/session"
)

// IndexController handles the GET request to get the index
func IndexController(ctx *gin.Context) {
	user := session.SessionsManager.GetUser(ctx)
	if user == nil {
		ctx.Redirect(302, "/login")
		return
	}
	ctx.HTML(200, "index.html", gin.H{
		"user": user,
	})
}
