// router.go

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/julianjjo/versasale-back/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
		public.GET("/user", controllers.User)
		public.GET("/product", controllers.Product)
		public.POST("/product", controllers.Product)
	}

	return router
}
