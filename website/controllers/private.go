package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user_name := c.PostForm("user_name")
		user_password := c.PostForm("user_password")
		var auth_user models.User
		err := utils.DB.Where("name = ?", user_name).First(&auth_user).Error
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"user_name_status": "is-invalid",
			})
			return
		}
		if utils.GetMD5Hash(user_password) != auth_user.Password {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"user_password_status": "is-invalid",
			})
			return
		}
		session.Set("user_name", auth_user.Name)
		session.Set("user_role", auth_user.Role)
		session.Set("user_active", auth_user.Active)
		if err := session.Save(); err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")

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
			log.Println("Failed to save session:", err)
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func DashboardPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user_name := session.Get("user_name")

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"button_text": "Logout",
			"button_link": "logout",
			"user_name":   user_name,
		})
	}
}
