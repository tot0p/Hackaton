package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackaton/model"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
	"strconv"
)

// TicketStatusPostController handles the POST request to change the status of a ticket
func TicketStatusPostController(ctx *gin.Context) {
	User := session.SessionsManager.GetUser(ctx)
	if User == nil {
		ctx.Redirect(302, "/login")
		return
	}

	if User.Role == model.RoleUser {
		ctx.Redirect(302, "/")
		return
	}

	uuid := ctx.Param("uuid")
	status := ctx.PostForm("status")

	if status == "" {
		ctx.Redirect(302, "/ticket/"+uuid+"?error=Invalid status")
		return
	}

	fmt.Println(status)

	statusInt, err := strconv.Atoi(status)
	if err != nil {
		ctx.Redirect(302, "/ticket/"+uuid+"?error=Invalid status")
		return
	}

	if statusInt < 0 || statusInt > 2 {
		ctx.Redirect(302, "/ticket/"+uuid+"?error=Invalid status")
		return
	}

	err = mysql.UpdateStatusTicket(uuid, statusInt)
	if err != nil {
		ctx.Redirect(302, "/ticket/"+uuid+"?error="+err.Error())
		return
	}

	ctx.Redirect(302, "/ticket/"+uuid)

}
