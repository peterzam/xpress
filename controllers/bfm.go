package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BfmItemsData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var bfm_items []models.Bfmitem

		if utils.DB.Find(&bfm_items).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		data, err := json.Marshal(bfm_items)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
