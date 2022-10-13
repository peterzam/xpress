package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
