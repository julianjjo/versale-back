package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/panic", func(c *gin.Context) {
		panic("Something went wrong!")
	})
	router.Run(":8080")
}
