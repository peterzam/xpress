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

	r.GET("/401", controllers.ErrorUnauthorizedPage())
	r.GET("/404", controllers.ErrorNotFoundPage())
	r.GET("/503", controllers.ErrorUnauthorizedPage())

	r.GET("/", controllers.StaticPages())
	r.GET("/aboutus", controllers.StaticPages())
	r.GET("/contactus", controllers.StaticPages())

	r.GET("/office", controllers.OfficePage())
	r.GET("/offices", controllers.OfficesData())

	r.GET("/login", controllers.StaticPages())
	r.POST("/login", controllers.LoginForm())

	r.GET("/register", controllers.StaticPages())
	r.POST("/register", controllers.RegisterForm())

	r.GET("/searchpackage", controllers.SearchPackagePage())
	r.POST("/searchpackage", controllers.SearchPackageForm())
}

func PrivateUserRoutes(r *gin.RouterGroup) {
	r.GET("/dashboard", controllers.DashboardPage())
	r.GET("/logout", controllers.Logout())

	r.GET("/addpackage", controllers.AddPackagePage())
	r.POST("/addpackage", controllers.AddPackageForm())

	r.GET("/request_pickup", controllers.RequestPickupPage())
	r.POST("/request_pickup", controllers.RequestPickupForm())

	r.GET("/addcomplaint", controllers.AddComplaintPage())
	r.POST("/addcomplaint", controllers.AddComplaintForm())
}

func PrivateAdminRoutes(r *gin.RouterGroup) {
	r.GET("/admin", controllers.AdminDashboardPage())
	r.GET("/report", controllers.ReportPage())
	r.GET("/report_data/:type/:value", controllers.ReportData())

	r.GET("/users", controllers.UsersData())
	r.GET("/manageuser", controllers.ManageUserPage())
	r.POST("/deleteuser", controllers.DeleteUserForm())
	r.POST("/edituser", controllers.EditUserForm())

	r.GET("/manageoffice", controllers.ManageOfficePage())
	r.POST("/addoffice", controllers.AddOfficeForm())
	r.POST("/deleteoffice", controllers.DeleteOfficeForm())
	r.POST("/editoffice", controllers.EditOfficeForm())

	r.GET("/packages", controllers.PackagesData())
	r.GET("/managepackage", controllers.ManagePackagePage())
	r.POST("/deletepackage", controllers.DeletePackageForm())
	r.POST("/editpackage", controllers.EditPackageForm())

	r.GET("/pickups", controllers.PickupsData())
	r.GET("/managepickup", controllers.ManagePickupPage())
	r.POST("/deletepickup", controllers.DeletePickupForm())

	r.GET("/complaints", controllers.ComplaintData())
	r.GET("/managecomplaint", controllers.ManageComplaintPage())
	r.POST("/deletecomplaint", controllers.DeleteComplaintsForm())
}
