package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"hackaton/controller"
	"hackaton/utils"
	"hackaton/utils/db/mongodb"
	"hackaton/utils/db/mysql"
)

var openBrowser *bool = new(bool)

func init() {
	err := env.Load()
	if err != nil {
		panic(err)
	}

	mysql.InitDB(env.Get("DB_HOST"), env.Get("DB_PORT"), env.Get("DB_USER"), env.Get("DB_PASSWORD"), env.Get("DB_DATABASE"))
	err = mongodb.NewMongoDB(env.Get("URI_MONGODB"))
	if err != nil {
		panic(err)
	}
	//gin.SetMode(gin.ReleaseMode)

	flag.BoolVar(openBrowser, "open", true, "Open browser automatically")
	flag.BoolVar(openBrowser, "o", true, "Open browser automatically")

	flag.Parse()

}

func main() {
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
	r.POST("/login", controller.LoginPostController)

	r.GET("/register", controller.RegisterController)
	r.POST("/register", controller.RegisterPostController)

	//API
	api := r.Group("/api")

	api.GET("/data/names", controller.DataSetNamesAPIController)
	api.GET("/data/:id", controller.DataSetByIdAPIController)

	if *openBrowser {
		utils.Openbrowser("http://localhost:" + env.Get("PORT"))
	}
	fmt.Println("Start server on port " + env.Get("PORT") + " ...")
	if err := r.Run(":" + env.Get("PORT")); err != nil {
		panic(err)
	}
}
