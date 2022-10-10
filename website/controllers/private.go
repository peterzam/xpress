package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1})
		if err := session.Save(); err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func ReportPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "report.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func ReportData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var chart models.Chart

		if c.Param("type") == "yearly" {
			value, _ := strconv.Atoi(c.Param("value"))
			chart = utils.YearlyReport(value)
		} else if c.Param("type") == "monthly" {
			value, _ := strconv.Atoi(c.Param("value"))
			fmt.Println("here")
			chart = utils.MonthlyReport(value)
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/404")
			return
		}

		data, err := json.Marshal(chart)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
