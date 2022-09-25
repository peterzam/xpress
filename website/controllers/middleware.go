package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user_id")
		if user == nil {
			log.Println("User not logged in")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
