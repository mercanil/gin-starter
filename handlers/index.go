package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeRouter(router *gin.Engine) {
	router.GET("/", showIndexPage)
}

func showIndexPage(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Anil Mercan Home Page",
		},
	)
}
