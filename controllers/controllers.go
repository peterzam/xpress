package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundPage(r *gin.Context) {
	r.HTML(http.StatusNotFound, "notfound.html", nil)
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
