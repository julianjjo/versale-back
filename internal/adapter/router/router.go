// router.go

package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	controllers "github.com/julianjjo/versasale-back/internal/adapter/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client, ctx context.Context) *gin.Engine {
	router := gin.Default()

	// Middleware to set client and ctx in Gin context
	router.Use(func(c *gin.Context) {
		c.Set("client", client)
		c.Next()
	})

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register-customer", controllers.RegisterCustomer)
		public.POST("/register-seller", controllers.RegisterSeller)
		public.POST("/login-customer", controllers.LoginCustomer)
		public.POST("/login-seller", controllers.LoginSeller)
		public.GET("/user", controllers.GetUser)
		public.GET("/product", controllers.GetProducts)
		public.GET("/product/:productId", controllers.GetProduct)
		public.POST("/product", controllers.CreateProduct)
	}

	return router
}
