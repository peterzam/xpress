package routes

import (
	"Xpress/controllers"
	"Xpress/utils"

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

	r.Use(sessions.Sessions("session", cookie.NewStore([]byte(utils.GetEnv("TOKEN_SECRET")))))

	// Routes
	/// No Routes
	r.NoRoute(controllers.NoRoute())
	r.NoMethod(controllers.NoRoute())
	/// Public Routes
	PublicRoutes(r.Group("/"))

	/// Private Routes
	private := r.Group("/")
	private.Use(controllers.AuthRequired())
	PrivateRoutes(private)

	return r
}

func PublicRoutes(r *gin.RouterGroup) {

	r.GET("/404", controllers.StaticPages())
	r.GET("/503", controllers.StaticPages())
	r.GET("/", controllers.StaticPages())
	r.GET("/aboutus", controllers.StaticPages())
	r.GET("/contactus", controllers.StaticPages())
	r.GET("/register", controllers.StaticPages())
	r.GET("/login", controllers.StaticPages())
	r.GET("/officemap", controllers.OfficeMapPage())
	r.GET("/offices", controllers.OfficesData())
	r.POST("/login", controllers.LoginForm())
	r.POST("/register", controllers.RegisterForm())
	r.GET("/searchpackage", controllers.SearchPackagePage())
	r.POST("/searchpackage", controllers.SearchPackageForm())
}

func PrivateRoutes(r *gin.RouterGroup) {
	r.GET("/dashboard", controllers.DashboardPage())
	r.GET("/report", controllers.ReportPage())
	r.GET("/logout", controllers.Logout())

	r.GET("/users", controllers.UsersData())
	r.GET("/manageuser", controllers.ManageUserPage())
	r.POST("/deleteuser", controllers.DeleteUserForm())
	r.POST("/edituser", controllers.EditUserForm())

	r.GET("/manageoffice", controllers.ManageOfficePage())
	r.POST("/addoffice", controllers.AddOfficeForm())
	r.POST("/deleteoffice", controllers.DeleteOfficeForm())
	r.POST("/editoffice", controllers.EditOfficeForm())

	r.GET("/addpackage", controllers.AddPackagePage())
	r.POST("/addpackage", controllers.AddPackageForm())

	r.GET("/packages", controllers.PackagesData())
	r.GET("/managepackage", controllers.ManagePackagePage())
	r.POST("/deletepackage", controllers.DeletePackageForm())
	r.POST("/editpackage", controllers.EditPackageForm())
}
