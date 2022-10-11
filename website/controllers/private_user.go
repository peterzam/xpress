package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var users []models.User

		if utils.DB.Find(&users).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}

func ManageUserPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "manageuser.html", gin.H{
			"button_text": "Admin Dashboard",
			"button_link": "admin",
		})
	}
}

func DeleteUserForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_phone := c.PostForm("user_phone")
		if utils.DB.Where("phone = ?", user_phone).Unscoped().Delete(models.User{}).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
	}
}

func EditUserForm() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth_user := models.User{
			Name:    c.PostForm("user_name"),
			Address: c.PostForm("user_address"),
			Role:    c.PostForm("user_role"),
		}

		if utils.DB.Model(models.User{}).Where("phone = ?", c.PostForm("user_phone")).Updates(auth_user).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
	}
}
