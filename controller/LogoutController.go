package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/session"
)

// LogoutController handles the GET request to logout
func LogoutController(ctx *gin.Context) {
	session.SessionsManager.DeleteSession(ctx)
	ctx.Redirect(302, "/login")
}
