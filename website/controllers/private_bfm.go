package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BfmPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "bfm.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "Dashboard",
		})
	}
}

func BfmItemsData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var bfm_items []models.Bfmitem

		if utils.DB.Find(&bfm_items).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		data, err := json.Marshal(bfm_items)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
