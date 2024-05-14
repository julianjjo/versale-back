// router.go

package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/julianjjo/versasale-back/internal/adapter/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
		public.GET("/user", controllers.GetUser)
		public.GET("/product", controllers.GetProducts)
		public.GET("/product/:productId", controllers.GetProduct)
		public.POST("/product", controllers.CreateProduct)
	}

	return router
}
