package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Err_NotFound() gin.HandlerFunc {
	return func(r *gin.Context) {
		r.HTML(http.StatusNotFound, "err_notfound.html", nil)
	}
}

func Err_Internal() gin.HandlerFunc {
	return func(r *gin.Context) {
		r.HTML(http.StatusInternalServerError, "err_internal.html", nil)
	}
}

func HomePage() gin.HandlerFunc {
	return func(r *gin.Context) {
		r.HTML(http.StatusOK, "home.html", nil)
	}
}

func AboutUsPage() gin.HandlerFunc {
	return func(r *gin.Context) {
		r.HTML(http.StatusOK, "aboutus.html", nil)
	}
}

func ContactUsPage() gin.HandlerFunc {
	return func(r *gin.Context) {
		r.HTML(http.StatusOK, "contactus.html", nil)
	}
}

func LoginPage() gin.HandlerFunc {
	return func(r *gin.Context) {
		r.HTML(http.StatusOK, "login.html", nil)
	}
}

func RegisterPage() gin.HandlerFunc {
	return func(r *gin.Context) {
		r.HTML(http.StatusOK, "register.html", nil)
	}
}
