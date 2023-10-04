package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
)

func NewTicketController(ctx *gin.Context) {
	if !session.SessionsManager.IsLogged(ctx) {
		ctx.Redirect(302, "/")
		return
	}
	ctx.HTML(200, "newticket.html", nil)
}

func NewTicketPostController(ctx *gin.Context) {
	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/login")
		return
	}

	Subject := ctx.PostForm("subject")
	Content := ctx.PostForm("content")
	LinkedData := ctx.PostForm("linked_data")

	if Subject == "" || Content == "" {
		ctx.HTML(200, "newticket.html", gin.H{
			"error": "Please fill Subject,Content the fields",
		})
		return
	}

	tk, err := mysql.CreateTicket(User.UUID, Subject, Content, LinkedData)
	if err != nil {
		ctx.HTML(200, "newticket.html", gin.H{
			"error": "An error occured while creating the ticket",
		})
		return
	}

	ctx.Redirect(302, "/ticket/"+tk.UUID)
}
