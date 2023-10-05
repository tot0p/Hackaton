package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackaton/model"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
)

func TicketController(ctx *gin.Context) {
	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/login")
		return
	}

	uuid := ctx.Param("uuid")

	tk, err := mysql.GetTicketByUUID(uuid)
	if err != nil {
		ctx.HTML(200, "ticket.html", gin.H{
			"error": err.Error(),
			"user":  User,
		})
		return
	}

	tchat := []*model.Tchat{}

	if tk.TchatUUID != "" {
		tchat, err = mysql.GetTchatByChannel(tk.TchatUUID)
		if err != nil {
			ctx.HTML(200, "ticket.html", gin.H{
				"error": err.Error(),
				"user":  User,
			})
			return
		}
	}

	tchatSorted := mysql.TchatToCompleteTchat(tchat)

	fmt.Println(tchat)

	fmt.Println(tchatSorted)

	if tk.UserUUID == User.UUID || User.Role != model.RoleUser {
		ctx.HTML(200, "ticket.html", gin.H{
			"ticket": tk,
			"tchat":  tchatSorted,
			"user":   User,
		})
		return
	}
	ctx.Redirect(302, "/")
}
