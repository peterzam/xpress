package controllers

import (
	"encoding/json"
	"net/http"
	"xpress/models"
	"xpress/utils"

	"github.com/gin-gonic/gin"
)

func UsersData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var users []models.User

		if utils.DB.Find(&users).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}

func DeleteUserForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_phone := c.PostForm("user_phone")
		if utils.DB.Where("phone = ?", user_phone).Unscoped().Delete(&models.User{}).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
	}
}

func EditUserForm() gin.HandlerFunc {
	return func(c *gin.Context) {

		var auth_user = models.User{
			Name:    c.PostForm("user_name"),
			Address: c.PostForm("user_address"),
			Role:    c.PostForm("user_role"),
		}

		if utils.DB.Where("phone = ?", c.PostForm("user_phone")).Updates(&auth_user).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
	}
}
