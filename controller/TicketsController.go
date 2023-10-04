package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/model"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
)

func TicketsController(ctx *gin.Context) {
	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/login")
		return
	}

	if User.Role == model.RoleUser {
		ctx.Redirect(302, "/")
		return
	}
	tickets, err := mysql.GetTicketsByUserUUID(User.UUID)
	if err != nil {
		ctx.HTML(200, "tickets.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(200, "tickets.html", gin.H{
		"tickets": tickets,
	})
}
