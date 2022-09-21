package routes

import (
	"Xpress/controllers"

	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	r := gin.Default()

	// Statics
	r.LoadHTMLGlob("static/templates/**/*")
	r.Static("assets", "static/assets")
	r.StaticFile("favicon.ico", "static/assets/favicon.ico")

	// Routes
	r.NoRoute(controllers.Err_NotFound)
	r.GET("/", controllers.HomePage)
	r.GET("/aboutus", controllers.AboutUsPage)
	r.GET("/contactus", controllers.ContactUsPage)
	r.GET("/login", controllers.LoginPage)
	r.GET("/register", controllers.RegisterPage)
	r.GET("/404", controllers.Err_NotFound)
	r.GET("/503", controllers.Err_Internal)

	r.POST("/login", controllers.LoginForm)

	return r
}
