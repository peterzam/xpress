package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PackagesData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var packages []models.Package

		if utils.DB.Find(&packages).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		data, err := json.Marshal(packages)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
func AddPackagePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "addpackage.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func AddPackageForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var p = &models.Package{
			Code:       utils.GeneratePackageCode(6),
			Dest_name:  c.PostForm("package_dest_name"),
			Dest_addr:  c.PostForm("package_dest_addr"),
			Dest_phone: c.PostForm("package_dest_phone"),
			Size:       c.PostForm("package_size"),
			Type:       c.PostForm("package_type"),
			Note:       c.PostForm("package_note"),
			Src_id:     session.Get("user_id").(uint),
			Created_at: time.Now().Unix(),
		}

		if utils.DB.Create(p).Error != nil {
			// c.Redirect(http.StatusTemporaryRedirect, "/503")

			c.HTML(http.StatusInternalServerError, "message.html", gin.H{
				"message_heading":     "Error 🤌",
				"message_text":        "Internal Server Error",
				"message_button":      "dashboard",
				"message_button_text": "Go back to Dashboard",
				"button_text":         "Dashboard",
				"button_link":         "dashboard",
			})
			return
		}
		var message []string
		message = append(message, "Tracking code : "+p.Code)
		message = append(message, " ")

		c.HTML(http.StatusOK, "message.html", gin.H{
			"message_heading":     "Success 👌",
			"message_text":        message,
			"message_button":      "dashboard",
			"message_button_text": "Go back to Dashboard",
			"button_text":         "Dashboard",
			"button_link":         "dashboard",
		})
	}
}

func ManagePackagePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "managepackage.html", gin.H{
			"button_text": "Admin Dashboard",
			"button_link": "admin",
		})
	}
}

func DeletePackageForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.DB.Where("code = ?", c.PostForm("package_code")).Unscoped().Delete(models.Package{}).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
	}
}

func EditPackageForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var auth_package = models.Package{
			Note:   c.PostForm("package_note"),
			Status: c.PostForm("package_status"),
		}
		if utils.DB.Model(models.Package{}).Where("code = ?", c.PostForm("package_code")).Updates(auth_package).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
	}
}
