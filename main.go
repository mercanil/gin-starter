package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Anil Mercan Home Page",
			},
		)
	})

	router.Run()
}
