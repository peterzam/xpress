package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"xpress/models"
	"xpress/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func ErrorPages() gin.HandlerFunc {
	return func(c *gin.Context) {
		status_code, _ := strconv.Atoi(c.Request.RequestURI[1:])

		var header = &gin.H{
			"button_text": "Login",
			"button_link": "login",
			"error_code":  status_code,
		}
		// Check session if already logged in
		session := sessions.Default(c)
		if session.Get("user_id") != nil {
			header = &gin.H{
				"button_text": "Dashboard",
				"button_link": "dashboard",
				"error_code":  status_code,
			}
		}

		c.HTML(status_code, "error.html", header)
	}
}

func StaticPages() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)

		// Login button
		var header = &gin.H{
			"button_text": "Login",
			"button_link": "login",
			"user_name":   session.Get("user_name"),
		}

		// Check session if already logged in
		switch session.Get("user_role") {
		case "user":
			header = &gin.H{
				"button_text": "Dashboard",
				"button_link": "dashboard",
				"user_name":   session.Get("user_name"),
			}

		case "admin":
			header = &gin.H{
				"button_text": "Admin Dashboard",
				"button_link": "admin",
				"user_name":   session.Get("user_name"),
			}
		}

		// Drop "?"
		uri := c.Request.RequestURI
		if uri[len(uri)-1:] == "?" {
			uri = uri[1 : len(uri)-1]
		} else {
			uri = uri[1:]
		}

		// Controlled route for "/" -> "home.html"
		if c.Request.RequestURI == "/" {
			c.HTML(http.StatusOK, "home.html", header)
		} else {
			c.HTML(http.StatusOK, fmt.Sprintf("%s.html", uri), header)
		}
	}
}

func LoginForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user_phone := c.PostForm("user_phone")
		user_password := c.PostForm("user_password")
		var auth_user models.User

		var invalid_login = gin.H{
			"button_text":       "Login",
			"button_link":       "login",
			"user_phone_status": "is-invalid",
		}

		if utils.DB.Where("phone = ?", user_phone).First(&auth_user).Error != nil {
			c.HTML(http.StatusOK, "login.html", invalid_login)
			return
		}
		if utils.GetMD5Hash(user_password) != auth_user.Password {
			c.HTML(http.StatusOK, "login.html", &invalid_login)
			return
		}
		session.Set("user_id", auth_user.Id)
		session.Set("user_phone", auth_user.Phone)
		session.Set("user_name", auth_user.Name)
		session.Set("user_address", auth_user.Address)
		session.Set("user_role", auth_user.Role)
		if err := session.Save(); err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1})
		if err := session.Save(); err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func RegisterForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = models.User{
			Name:     c.PostForm("user_name"),
			Phone:    c.PostForm("user_phone"),
			Address:  c.PostForm("user_address"),
			Password: c.PostForm("user_password"),
		}

		// Check User Inputs
		check := utils.CheckUserInputs(user)
		if check != "" {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"user_register_status_text": check,
			})
			return
		}

		user.Password = utils.GetMD5Hash(user.Password)
		if utils.DB.Create(&user).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func SearchPackageForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pack = models.Package{
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
			c.Redirect(http.StatusTemporaryRedirect, "/500")
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
			"message_button_text": "Go back",
			"button_text":         "Dashboard",
			"button_link":         "dashboard",
		})

	}
}

// func PackageData() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		type Message struct {
// 			Package        models.Package
// 			Sender_Name    string `json:"sender_name"`
// 			Sender_Address string `json:"sender_address"`
// 		}
// 		var pack = models.Package{
// 			Code: c.Query("code"),
// 		}
// 		if utils.DB.Where("code = ?", pack.Code).First(&pack).Error != nil {
// 			c.Redirect(http.StatusTemporaryRedirect, "/404")
// 			return
// 		}
// 		var src_user models.User
// 		if utils.DB.Where("id = ?", pack.Src_id).First(&src_user).Error != nil {
// 			c.Redirect(http.StatusTemporaryRedirect, "/500")
// 			return
// 		}
// 		data, err := json.Marshal(&Message{
// 			Package:        pack,
// 			Sender_Name:    src_user.Name,
// 			Sender_Address: src_user.Address,
// 		})
// 		if err != nil {
// 			c.Redirect(http.StatusTemporaryRedirect, "/500")
// 			return
// 		}
// 		c.Data(http.StatusOK, gin.MIMEJSON, data)

// 	}
// }

func OfficesData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var offices []models.Office

		if utils.DB.Find(&offices).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		data, err := json.Marshal(offices)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
