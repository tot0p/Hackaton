package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/session"
)

func ProfilController(ctx *gin.Context) {
	if !session.SessionsManager.IsLogged(ctx) {
		ctx.Redirect(302, "/login")
		return
	}

	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/login")
		return
	}
	ctx.HTML(200, "profil.html", gin.H{
		"user": User,
	})
}
