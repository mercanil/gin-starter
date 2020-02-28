package main

import (
	"../handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	handlers.InitializeRouter(router)
	router.Run()
}
