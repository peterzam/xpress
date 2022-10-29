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

	/// Private - user Routes
	user := r.Group("/")
	user.Use(controllers.UserAuthRequired())
	PrivateUserRoutes(user)

	/// Private - admin Routes
	admin := r.Group("/")
	admin.Use(controllers.AdminAuthRequired())
	PrivateAdminRoutes(admin)

	return r
}

func PublicRoutes(r *gin.RouterGroup) {

	r.GET("/401", controllers.ErrorPages())
	r.GET("/404", controllers.ErrorPages())
	r.GET("/500", controllers.ErrorPages())

	r.GET("/", controllers.StaticPages())
	r.GET("/aboutus", controllers.StaticPages())
	r.GET("/contactus", controllers.StaticPages())

	r.GET("/office", controllers.StaticPages())
	r.GET("/offices", controllers.OfficesData())

	r.GET("/login", controllers.StaticPages())
	r.POST("/login", controllers.LoginForm())

	r.GET("/register", controllers.StaticPages())
	r.POST("/register", controllers.RegisterForm())

	r.GET("/searchpackage", controllers.StaticPages())
	r.POST("/searchpackage", controllers.SearchPackageForm())
	// r.GET("/package", controllers.PackageData())
	// r.GET("/showpackage", controllers.StaticPages())
}

func PrivateUserRoutes(r *gin.RouterGroup) {
	r.GET("/dashboard", controllers.StaticPages())
	r.GET("/logout", controllers.Logout())

	r.GET("/addpackage", controllers.StaticPages())
	r.POST("/addpackage", controllers.AddPackageForm())

	r.GET("/addpickup", controllers.StaticPages())
	r.POST("/addpickup", controllers.AddPickupForm())

	r.GET("/addcomplaint", controllers.StaticPages())
	r.POST("/addcomplaint", controllers.AddComplaintForm())

	r.GET("/addbfm", controllers.StaticPages())
	r.GET("/bfms", controllers.BfmItemsData())
	r.POST("/addbfm", controllers.AddBfmForm())
	r.GET("/userpackages", controllers.UserPackagesData())
	r.GET("/showpackages", controllers.StaticPages())
}

func PrivateAdminRoutes(r *gin.RouterGroup) {
	r.GET("/admin", controllers.StaticPages())

	r.GET("/report", controllers.StaticPages())
	r.GET("/report_data/:type/:value", controllers.ReportData())

	r.GET("/users", controllers.UsersData())
	r.GET("/manageuser", controllers.StaticPages())
	r.POST("/deleteuser", controllers.DeleteUserForm())
	r.POST("/edituser", controllers.EditUserForm())

	r.GET("/manageoffice", controllers.StaticPages())
	r.POST("/addoffice", controllers.AddOfficeForm())
	r.POST("/deleteoffice", controllers.DeleteOfficeForm())
	r.POST("/editoffice", controllers.EditOfficeForm())

	r.GET("/allpackages", controllers.AllPackagesData())
	r.GET("/managepackage", controllers.StaticPages())
	r.POST("/deletepackage", controllers.DeletePackageForm())
	r.POST("/editpackage", controllers.EditPackageForm())

	r.GET("/pickups", controllers.PickupsData())
	r.GET("/managepickup", controllers.StaticPages())
	r.POST("/deletepickup", controllers.DeletePickupForm())

	r.GET("/complaints", controllers.ComplaintData())
	r.GET("/managecomplaint", controllers.StaticPages())
	r.POST("/deletecomplaint", controllers.DeleteComplaintsForm())
}
