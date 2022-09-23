package routes

import (
	"Xpress/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	r := gin.Default()

	// Statics
	r.LoadHTMLGlob("static/templates/**/*")
	r.Static("assets", "static/assets")
	r.StaticFile("favicon.ico", "static/assets/favicon.ico")

	r.Use(sessions.Sessions("session", cookie.NewStore([]byte("globals.Secret"))))

	// Routes
	/// No Routes
	r.NoRoute(controllers.Err_NotFound())

	/// Public Routes
	PublicRoutes(r.Group("/"))

	/// Public Routes
	private := r.Group("/")
	private.Use(controllers.AuthRequired())
	PrivateRoutes(private)

	return r
}

func PublicRoutes(r *gin.RouterGroup) {

	r.GET("/404", controllers.Err_NotFound())
	r.GET("/503", controllers.Err_Internal())

	r.GET("/", controllers.HomePage())
	r.GET("/aboutus", controllers.AboutUsPage())
	r.GET("/contactus", controllers.ContactUsPage())
	r.GET("/register", controllers.RegisterPage())

	r.GET("/login", controllers.LoginPage())
	r.POST("/login", controllers.LoginForm())
}

func PrivateRoutes(r *gin.RouterGroup) {
	r.GET("/dashboard", controllers.DashboardPage())
	r.GET("/logout", controllers.Logout())
}
