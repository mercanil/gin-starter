package main

import (
	"../route"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	route.InitializeRouter(router)
	router.Run()
}
