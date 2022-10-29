package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddOfficeForm() gin.HandlerFunc {
	return func(c *gin.Context) {

		var auth_office = models.Office{
			Name:         c.PostForm("office_name"),
			Address:      c.PostForm("office_address"),
			Phone:        c.PostForm("office_phone"),
			Location_lat: c.PostForm("office_location_lat"),
			Location_lng: c.PostForm("office_location_lng"),
		}
		if utils.DB.Create(&auth_office).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
	}
}

func DeleteOfficeForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.DB.Where("id = ?", c.PostForm("office_id")).Unscoped().Delete(&models.Office{}).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
	}
}

func EditOfficeForm() gin.HandlerFunc {
	return func(c *gin.Context) {

		var auth_office = models.Office{
			Name:         c.PostForm("office_name"),
			Address:      c.PostForm("office_address"),
			Phone:        c.PostForm("office_phone"),
			Location_lat: c.PostForm("office_location_lat"),
			Location_lng: c.PostForm("office_location_lng"),
		}
		if utils.DB.Where("id = ?", c.PostForm("office_id")).Updates(&auth_office).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
	}
}
