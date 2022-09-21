package controllers

import (
	"Xpress/models"
	"Xpress/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// var user models.User

func LoginForm(r *gin.Context) {
	user_name := r.PostForm("user_name")
	user_password := r.PostForm("user_password")
	var auth_user models.User
	err := utils.DB.Where("name = ?", user_name).First(&auth_user).Error
	if err != nil {
		r.HTML(http.StatusOK, "login.html", gin.H{
			"user_name_status": "is-invalid",
		})
		return
	}
	if utils.GetMD5Hash(user_password) != auth_user.Password {
		r.HTML(http.StatusOK, "login.html", gin.H{
			"user_password_status": "is-invalid",
		})
		return
	}
	token, err := utils.GenerateJWT(auth_user)
	if err != nil {
		r.Redirect(http.StatusTemporaryRedirect, "/503")
		return
	}

	token_lifetime, _ := strconv.Atoi(utils.GetEnv("TOKEN_LIFETIME"))
	expire_at := int(time.Now().Add(time.Duration(token_lifetime) * time.Minute).Unix())
	token_secure, _ := strconv.ParseBool(utils.GetEnv("TOKEN_SECURE"))
	token_http_only, _ := strconv.ParseBool(utils.GetEnv("TOKEN_HTTP_ONLY"))

	r.SetCookie(
		"session",
		token, expire_at, "/",
		utils.GetEnv("WEB_ADDRESS"),
		token_secure,
		token_http_only,
	)
	result := fmt.Sprintf("{ user_id : %d, user_name : %s, user_role : %s , user_active : %t}",
		auth_user.Id, auth_user.Name, auth_user.Role, auth_user.Active)
	r.JSON(http.StatusOK, result)
}
