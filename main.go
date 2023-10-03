package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"hackaton/controller"
)

func init() {
	err := env.Load()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("src/templates/*")

	r.Static("/html", "./src/html")
	r.Static("/css", "./src/css")
	r.Static("/js", "./src/js")

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
	api.POST("/register", controller.RegisterAPIController)

	if err := r.Run(":" + env.Get("PORT")); err != nil {
		panic(err)
	}
}
