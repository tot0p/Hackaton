package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/db/mysql"
	"hackaton/utils/session"
)

// LoginController handles the GET request to login
func LoginController(ctx *gin.Context) {
	if session.SessionsManager.IsLogged(ctx) {
		ctx.Redirect(302, "/")
		return
	}
	ctx.HTML(200, "login.html", nil)
}

// LoginPostController handles the POST request to login
func LoginPostController(ctx *gin.Context) {
	if session.SessionsManager.IsLogged(ctx) {
		ctx.Redirect(302, "/")
		return
	}
	Pseudo := ctx.PostForm("pseudo")
	Password := ctx.PostForm("password")

	if Pseudo == "" || Password == "" {
		ctx.HTML(200, "login.html", gin.H{
			"error": "Please fill all fields",
		})
		return
	}

	User, err := mysql.GetUserByUsername(Pseudo)
	if err != nil {
		ctx.HTML(200, "login.html", gin.H{
			"error": "user not found",
		})
		return
	}

	if !User.ComparePassword(Password) {
		ctx.HTML(200, "login.html", gin.H{
			"error": "Wrong password",
		})
		return
	}

	User.RemovePassword()

	session.SessionsManager.CreateSession(ctx, &User)

	ctx.Redirect(302, "/")
}
