package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"xpress/models"
	"xpress/utils"

	"github.com/gin-contrib/sessions"
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

func AddBfmForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		var bfmitem models.Bfmitem
		if utils.DB.Where("id = ?", c.PostForm("bfmitem_id")).First(&bfmitem).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
			return
		}

		var p = models.Package{
			Code:       utils.GeneratePackageCode(6),
			Dest_name:  session.Get("user_name").(string),
			Dest_addr:  session.Get("user_address").(string),
			Dest_phone: session.Get("user_phone").(string),
			Size:       bfmitem.Size,
			Type:       "bfmitem",
			Note:       "bfmitem_" + bfmitem.Name + c.PostForm("bfmitem_user_note"),
			Src_id:     session.Get("user_id").(uint),
			Created_at: time.Now().Unix(),
		}

		if utils.DB.Create(&p).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/500")
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
