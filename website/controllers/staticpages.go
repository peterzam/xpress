package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Err_NotFound(r *gin.Context) {
	r.HTML(http.StatusNotFound, "err_notfound.html", nil)
}

func Err_Internal(r *gin.Context) {
	r.HTML(http.StatusInternalServerError, "err_internal.html", nil)
}

func HomePage(r *gin.Context) {
	r.HTML(http.StatusOK, "home.html", nil)
}

func AboutUsPage(r *gin.Context) {
	r.HTML(http.StatusOK, "aboutus.html", nil)
}

func ContactUsPage(r *gin.Context) {
	r.HTML(http.StatusOK, "contactus.html", nil)
}

func LoginPage(r *gin.Context) {
	r.HTML(http.StatusOK, "login.html", nil)
}

func RegisterPage(r *gin.Context) {
	r.HTML(http.StatusOK, "register.html", nil)
}
