package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Err_NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/404")
	}
}

func StaticPages() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Login button
		var header = gin.H{
			"button_text": "Login",
			"button_link": "login",
		}
		// Check session if already logged in
		session := sessions.Default(c)
		if session.Get("user_name") != nil {
			header = gin.H{
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
