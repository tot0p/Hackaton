package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
)

func TchatPostController(ctx *gin.Context) {
	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/login")
		return
	}

	uuidTchat := ctx.Param("uuidTchat")

	msg := ctx.PostForm("msg")

	author := User.UUID

	channel := ctx.PostForm("channel")

	if channel == "" {
		ctx.Redirect(302, "/ticket/"+uuidTchat+"?error=Channel vide")
		return
	}

	if msg == "" {
		ctx.Redirect(302, "/ticket/"+uuidTchat+"?error=Message vide")
		return
	}

	_, err := mysql.CreateTchat(author, msg, channel)
	if err != nil {
		ctx.Redirect(302, "/ticket/"+uuidTchat+"?error="+err.Error())
		return
	}

	ctx.Redirect(302, "/ticket/"+uuidTchat)

}
