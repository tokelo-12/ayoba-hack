package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//auth
	// Login()

	router := gin.Default()

	router.POST("/sendtext", sendMsg)

	router.GET("/gettext", getText)

	router.POST("/createcard", createCard)

	router.Run("localhost:8080")

}
