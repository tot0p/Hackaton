package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"hackaton/controller"
	"hackaton/utils/db/mysql"
)

func init() {
	err := env.Load()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)

}

func main() {
	mysql.InitDB(env.Get("DB_HOST"), env.Get("DB_PORT"), env.Get("DB_USER"), env.Get("DB_PASSWORD"), env.Get("DB_DATABASE"))
	r := gin.Default()
	r.LoadHTMLGlob("src/templates/*")

	r.Static("/static", "./src/static")

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
	api.POST("/data/:id", controller.DataSetByIdAPIController)

	fmt.Println("Start server on port " + env.Get("PORT") + " ...")
	if err := r.Run(":" + env.Get("PORT")); err != nil {
		panic(err)
	}
}
