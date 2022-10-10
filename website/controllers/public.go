package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func StaticPages() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Login button
		var header = &gin.H{
			"button_text": "Login",
			"button_link": "login",
		}
		// Check session if already logged in
		session := sessions.Default(c)
		if session.Get("user_id") != nil {
			header = &gin.H{
				"button_text": "Dashboard",
				"button_link": "dashboard",
			}
		}

		// Controlled route for "/" -> "home.html"
		if c.Request.RequestURI == "/" {
			c.HTML(http.StatusOK, "home.html", header)
		} else {
			c.HTML(http.StatusOK, fmt.Sprintf("%s.html", c.Request.RequestURI[1:]), header)
		}
	}
}

func LoginForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user_phone := c.PostForm("user_phone")
		user_password := c.PostForm("user_password")
		var auth_user models.User

		var invalid_login = &gin.H{
			"button_text":       "Login",
			"button_link":       "login",
			"user_phone_status": "is-invalid",
		}

		if utils.DB.Where("phone = ?", user_phone).First(&auth_user).Error != nil {
			c.HTML(http.StatusOK, "login.html", invalid_login)
			return
		}
		if utils.GetMD5Hash(user_password) != auth_user.Password {
			c.HTML(http.StatusOK, "login.html", invalid_login)
			return
		}
		session.Set("user_id", auth_user.Id)
		session.Set("user_phone", auth_user.Phone)
		session.Set("user_name", auth_user.Name)
		session.Set("user_address", auth_user.Address)
		session.Set("user_role", auth_user.Role)
		if err := session.Save(); err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func RegisterForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{
			Name:     c.PostForm("user_name"),
			Phone:    c.PostForm("user_phone"),
			Address:  c.PostForm("user_address"),
			Password: c.PostForm("user_password"),
		}

		// Check User Inputs
		check := utils.CheckUserInputs(*user)
		if check != "" {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"user_register_status_text": check,
			})
			return
		}

		user.Password = utils.GetMD5Hash(user.Password)
		if utils.DB.Create(user).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func SearchPackagePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "searchpackage.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func SearchPackageForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pack = &models.Package{
			Code: c.PostForm("package_code"),
		}
		if utils.DB.Where("code = ?", pack.Code).First(&pack).Error != nil {
			c.HTML(http.StatusOK, "searchpackage.html", gin.H{
				"package_search_response_text": "Package not found",
			})
			return
		}

		var src_user models.User
		if utils.DB.Where("id = ?", pack.Src_id).First(&src_user).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		var message []string
		message = append(message, "Sender Name : "+src_user.Name)
		message = append(message, "Sender Phone : "+src_user.Phone)
		message = append(message, "Sender Address : "+src_user.Address)
		message = append(message, "Reveiver Name : "+pack.Dest_name)
		message = append(message, "Reveiver Phone : "+pack.Dest_phone)
		message = append(message, "Reveiver Address : "+pack.Dest_addr)
		message = append(message, "Package Size : "+pack.Size)
		message = append(message, "Package Type : "+pack.Type)
		message = append(message, "Package Note : "+pack.Note)
		message = append(message, "Package Status : "+pack.Status)
		message = append(message, "Send Time : "+time.Unix(pack.Created_at, 0).Format(time.UnixDate))

		c.HTML(http.StatusOK, "message.html", gin.H{
			"message_heading":     "Package 🎁",
			"message_text":        message,
			"message_button":      "dashboard",
			"message_button_text": "Go back to Dashboard",
			"button_text":         "Dashboard",
			"button_link":         "dashboard",
		})

	}
}

func OfficeMapPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "office.html", gin.H{
			"button_text": "Login",
			"button_link": "login",
			"map_api_key": utils.GetEnv("MAP_API_KEY"),
		})
	}
}

func OfficesData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var offices []models.Office

		if utils.DB.Find(&offices).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		data, err := json.Marshal(offices)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
