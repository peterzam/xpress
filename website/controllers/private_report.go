package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
