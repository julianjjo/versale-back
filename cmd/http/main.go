package main

import (
	"github.com/gin-gonic/gin"
	routers "github.com/julianjjo/versasale-back/internal/adapter/router"
)

func main() {
	router := gin.Default()

	router = routers.SetupRouter()
	router.Use(gin.Recovery())

	router.GET("/panic", func(c *gin.Context) {
		panic("Something went wrong!")
	})
	router.Run(":8080")
}
