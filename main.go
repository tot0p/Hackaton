package main

import (
	"github.com/gin-gonic/gin"
	"hackaton/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("src/templates/*")

	r.Static("/html", "./src/html")
	r.Static("/css", "./src/css")
	r.Static("/js", "./src/js")

	//gin.SetMode(gin.ReleaseMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Index
	r.GET("/", controller.IndexController)
	r.GET("/index", controller.IndexController)

	r.GET("/login", controller.LoginController)

	//API
	api := r.Group("/api")
	api.POST("/login", controller.LoginAPIController)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
