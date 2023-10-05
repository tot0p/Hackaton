package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/db/mysql"
	"strconv"
)

func TicketStatusPostController(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	status := ctx.PostForm("status")

	if status == "" {
		ctx.Redirect(302, "/ticket/"+uuid+"?error=Invalid status")
		return
	}

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
