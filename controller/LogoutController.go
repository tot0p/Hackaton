package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/session"
)

func LogoutController(ctx *gin.Context) {
	session.SessionsManager.DeleteSession(ctx)
	ctx.Redirect(302, "/login")
}
