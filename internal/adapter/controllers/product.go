package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": []gin.H{
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 1",
				"price":         1000,
				"stock":         10,
				"description":   "Product 1 description",
				"product_image": "https://via.placeholder.com/200",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 2",
				"price":         2000,
				"stock":         20,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 2 description",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 3",
				"price":         3000,
				"stock":         30,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 3 description",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 4",
				"price":         4000,
				"stock":         40,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 4 description",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 5",
				"price":         5000,
				"stock":         50,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 5 description",
			},
		},
	},
	)
}

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "test"})
}
