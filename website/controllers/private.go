package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DashboardPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user_name := session.Get("user_name")

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"button_text": "Logout",
			"button_link": "logout",
			"user_name":   user_name,
		})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{Path: "/", MaxAge: -1})
		if err := session.Save(); err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
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
			Note:       c.PostForm("package_dest_note"),
			Src_id:     session.Get("user_id").(uint),
			Created_at: time.Now().Unix(),
		}
		fmt.Println(p)

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

func SearchPackagePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "searchpackage.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func SearchPackageForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pack = &models.Package{
			Code: c.PostForm("package_code"),
		}
		if utils.DB.Where("code = ?", pack.Code).First(&pack).Error != nil {
			c.HTML(http.StatusOK, "searchpackage.html", gin.H{
				"package_search_response_text": "Package not found",
			})
			return
		}

		var src_user models.User
		if utils.DB.Where("id = ?", pack.Src_id).First(&src_user).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		var message []string
		message = append(message, "Sender Name : "+src_user.Name)
		message = append(message, "Sender Phone : "+src_user.Phone)
		message = append(message, "Sender Address : "+src_user.Address)
		message = append(message, "Reveiver Name : "+pack.Dest_name)
		message = append(message, "Reveiver Phone : "+pack.Dest_phone)
		message = append(message, "Reveiver Address : "+pack.Dest_addr)
		message = append(message, "Package Size : "+pack.Size)
		message = append(message, "Package Type : "+pack.Type)
		message = append(message, "Package Note : "+pack.Note)
		message = append(message, "Receive Status : "+strconv.FormatBool(pack.Active))
		message = append(message, "Send Time : "+time.Unix(pack.Created_at, 0).Format(time.UnixDate))

		c.HTML(http.StatusOK, "message.html", gin.H{
			"message_heading":     "Package 🎁",
			"message_text":        message,
			"message_button":      "dashboard",
			"message_button_text": "Go back to Dashboard",
			"button_text":         "Dashboard",
			"button_link":         "dashboard",
		})

	}
}

func ReportPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "report.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func ManageUserPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "manageuser.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func EditUserPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "edituser.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func DeleteUserPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "deleteuser.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
		})
	}
}

func OfficeMapPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "officemap.html", gin.H{
			"button_text": "Dashboard",
			"button_link": "dashboard",
			"map_api_key": utils.GetEnv("MAP_API_KEY"),
		})
	}
}

func OfficeLocationsData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var offices []models.Office

		if utils.DB.Find(&offices).Error != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		data, err := json.Marshal(offices)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/503")
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}

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

func DeleteUserForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_phone := c.PostForm("user_phone")
		var auth_user models.User
		if utils.DB.Where("phone = ?", user_phone).Unscoped().Delete(&auth_user).Error != nil {
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
