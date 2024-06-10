package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julianjjo/versasale-back/internal/infrastructure/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func DocumentType(c *gin.Context) {
	c.Set("status", "success")

	client, _ := c.MustGet("client").(*mongo.Client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	documentTypes, err := service.GetDocumentType(client, ctx)
	if err != nil {
		c.Set("status", "error") // Set status to "error"
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": c.GetString("status"), // Use the status from the context
		"data":   documentTypes,
	})
}
