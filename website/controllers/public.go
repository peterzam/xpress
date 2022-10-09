package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"fmt"
	"net/http"

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
