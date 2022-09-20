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
	r.NoRoute(controllers.NotFoundPage)

	r.GET("/", controllers.HomePage)
	r.GET("/aboutus", controllers.AboutUsPage)
	r.GET("/contactus", controllers.ContactUsPage)
	r.GET("/login", controllers.LoginPage)
	r.GET("/register", controllers.RegisterPage)

	return r
}
