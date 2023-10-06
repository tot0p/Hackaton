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
	// Load env file
	err := env.Load()
	if err != nil {
		panic(err)
	}

	// Init DB
	mysql.InitDB(env.Get("DB_HOST"), env.Get("DB_PORT"), env.Get("DB_USER"), env.Get("DB_PASSWORD"), env.Get("DB_DATABASE"))
	err = mongodb.NewMongoDB(env.Get("URI_MONGODB"))
	if err != nil {
		panic(err)
	}

	//Release mode
	gin.SetMode(gin.ReleaseMode)

	// Open browser flag
	flag.BoolVar(openBrowser, "open", false, "Open browser automatically")
	flag.BoolVar(openBrowser, "o", false, "Open browser automatically")
	flag.Parse()

}

func main() {
	r := gin.Default()
	// Load templates
	r.LoadHTMLGlob("src/templates/*")

	// Serve static files
	r.Static("/static", "./src/static")

	// Index
	r.GET("/", controller.IndexController)
	r.GET("/index", controller.IndexController)

	//Login
	r.GET("/login", controller.LoginController)
	r.POST("/login", controller.LoginPostController)

	//Logout
	r.GET("/logout", controller.LogoutController)

	//Register
	r.GET("/register", controller.RegisterController)
	r.POST("/register", controller.RegisterPostController)

	//Ticket
	r.GET("/newticket", controller.NewTicketController)
	r.POST("/newticket", controller.NewTicketPostController)
	r.GET("/ticket/:uuid", controller.TicketController)
	r.GET("/tickets", controller.TicketsController)
	r.POST("/ticket/:uuid/status", controller.TicketStatusPostController)

	//Tchat
	r.POST("/tchat/:uuidticket", controller.TchatPostController)

	//API Group
	api := r.Group("/api")

	// DataSet
	api.GET("/data/names", controller.DataSetNamesAPIController)
	api.GET("/data/category", controller.DataSetCategoriesAPIController)
	api.GET("/data/category/:category", controller.DataSetByCategoryAPIController)
	api.GET("/data/name/:name", controller.DataSetByNameAPIController)
	api.GET("/data/name/:name/fields", controller.DataSetByNameGetFieldsAPIController)
	api.GET("/data/name/:name/filter", controller.DataSetByNameFilterAPIController)
	api.GET("/data/name/:name/filters", controller.DataSetByNameFiltersAPIController)
	api.GET("/data/name/:name/field/:field/value/:value", controller.DataSetByNameFieldAPIController)

	// Graph
	api.GET("/graph", controller.GraphAPIController)
	api.GET("/graph/:name", controller.GraphByNameAPIController)

	// Open browser if flag is true
	if *openBrowser {
		utils.OpenBrowser("http://localhost:" + env.Get("PORT"))
	}

	// Launch server
	fmt.Println("Start server on port " + env.Get("PORT") + " ...")
	if err := r.Run(":" + env.Get("PORT")); err != nil {
		panic(err)
	}
}
