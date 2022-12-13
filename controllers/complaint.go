package controllers

import (
	"encoding/json"
	"net/http"
	"xpress/models"
	"xpress/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AddComplaintForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var p = models.Complaint{
			Note:   c.PostForm("note"),
			Src_id: session.Get("user_id").(uint),
		}

		if utils.DB.Create(&p).Error != nil {

			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}

		c.HTML(http.StatusOK, "message.html", gin.H{
			"message_heading":     "Complaint Received 🫰",
			"message_button":      "dashboard",
			"message_button_text": "Go back to Dashboard",
			"button_text":         "Dashboard",
			"button_link":         "dashboard",
		})
	}
}

func ComplaintData() gin.HandlerFunc {
	return func(c *gin.Context) {

		type Response struct {
			Id         uint   `json:"complaint_id"`
			Note       string `json:"complaint_note"`
			User_name  string `json:"user_name"`
			User_phone string `json:"user_phone"`
		}

		var complaints []models.Complaint
		var response []Response

		if utils.DB.Find(&complaints).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}

		for _, complaint := range complaints {
			var src_user models.User
			if utils.DB.Where("id = ?", complaint.Src_id).First(&src_user).Error != nil {
				c.Redirect(http.StatusTemporaryRedirect, "/500")
				return
			}
			response = append(response, Response{
				Id:         complaint.Id,
				Note:       complaint.Note,
				User_name:  src_user.Name,
				User_phone: src_user.Phone,
			})
		}

		data, err := json.Marshal(response)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}

func DeleteComplaintsForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.DB.Where("id = ?", c.PostForm("complaint_id")).Unscoped().Delete(&models.Complaint{}).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}
	}
}
