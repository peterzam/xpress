package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequestPickupPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "requestpickup.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func RequestPickupForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var p = &models.Pickup{
			Size:   c.PostForm("pickup_size"),
			Type:   c.PostForm("pickup_type"),
			Note:   c.PostForm("pickup_note"),
			Src_id: session.Get("user_id").(uint),
		}
		fmt.Println(p)

		if utils.DB.Create(p).Error != nil {

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

		c.HTML(http.StatusOK, "message.html", gin.H{
			"message_heading":     "Success 👌",
			"message_button":      "dashboard",
			"message_button_text": "Go back to Dashboard",
			"button_text":         "Dashboard",
			"button_link":         "dashboard",
		})
	}
}
func PickupsData() gin.HandlerFunc {
	return func(c *gin.Context) {

		type Response struct {
			Id         uint   `json:"pickup_id"`
			Note       string `json:"pickup_note"`
			User_name  string `json:"user_name"`
			User_phone string `json:"user_phone"`
			User_addr  string `json:"user_address"`
		}

		var pickups []models.Pickup
		var response []Response

		if utils.DB.Find(&pickups).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}

		for _, pickup := range pickups {
			var src_user models.User
			if utils.DB.Where("id = ?", pickup.Src_id).First(&src_user).Error != nil {
				c.Redirect(http.StatusTemporaryRedirect, "/503")
				return
			}
			var resp Response
			resp.Id = pickup.Id
			resp.Note = pickup.Note
			resp.User_name = src_user.Name
			resp.User_phone = src_user.Phone
			resp.User_addr = src_user.Address
			response = append(response, resp)
		}

		data, err := json.Marshal(response)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}

func ManagePickupPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "managepickup.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func DeletePickupForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.DB.Where("id = ?", c.PostForm("pickup_id")).Unscoped().Delete(models.Pickup{}).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
	}
}
