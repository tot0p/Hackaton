package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/model"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
)

func RegisterController(ctx *gin.Context) {
	ctx.HTML(200, "register.html", nil)
}

func RegisterPostController(ctx *gin.Context) {
	Pseudo := ctx.PostForm("pseudo")
	Password := ctx.PostForm("password")
	Email := ctx.PostForm("email")
	Phone := ctx.PostForm("phone")
	Role := model.RoleUser
	/*
		if Pseudo == "" || Password == "" || Email == "" || Phone == "" {
			ctx.HTML(200, "register.html", gin.H{
				"error": "Please fill all fields",
			})
			return
		} else if !FormValidatorData.IsValidFrenchPhoneNumber(Phone) {
			ctx.HTML(200, "register.html", gin.H{
				"error": "Phone number is not valid",
			})
			return
		} else if !FormValidatorData.IsValidEmail(Email) {
			ctx.HTML(200, "register.html", gin.H{
				"error": "Email is not valid",
			})
			return
		} else if !FormValidatorData.IsPasswordSecure(Password) {
			ctx.HTML(200, "register.html", gin.H{
				"error": "Password is not secure",
			})
			return
		} else if !FormValidatorData.IsValidPseudo(Pseudo) {
			ctx.HTML(200, "register.html", gin.H{
				"error": "Pseudo is not valid",
			})
			return
		}
	*/

	User, err := mysql.CreateUser(Pseudo, Password, Email, Phone, Role)
	User.RemovePassword()
	if err != nil {
		ctx.HTML(200, "register.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	session.SessionsManager.CreateSession(ctx, &User)

	ctx.Redirect(302, "/")
}
