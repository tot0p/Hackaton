package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/model"
	"hackaton/utils/db/mysql"
)

func RegisterAPIController(ctx *gin.Context) {
	Pseudo := ctx.PostForm("pseudo")
	Password := ctx.PostForm("password")
	Email := ctx.PostForm("email")
	Phone := ctx.PostForm("phone")
	Role := model.RoleUser

	User, err := mysql.CreateUser(Pseudo, Password, Email, Phone, Role)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, User)
}
