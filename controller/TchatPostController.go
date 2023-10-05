package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
)

// TchatPostController handles the POST request to send a message in a ticket
func TchatPostController(ctx *gin.Context) {
	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/login")
		return
	}

	uuidTicket := ctx.Param("uuidticket")

	msg := ctx.PostForm("msg")

	author := User.UUID

	channel := ctx.PostForm("channel")

	if channel == "" {
		ctx.Redirect(302, "/ticket/"+uuidTicket+"?error=Channel vide")
		return
	}

	if msg == "" {
		ctx.Redirect(302, "/ticket/"+uuidTicket+"?error=Message vide")
		return
	}

	_, err := mysql.CreateTchat(author, msg, channel)
	if err != nil {
		ctx.Redirect(302, "/ticket/"+uuidTicket+"?error="+err.Error())
		return
	}

	ctx.Redirect(302, "/ticket/"+uuidTicket)

}
