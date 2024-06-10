package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/julianjjo/versasale-back/internal/core/model"
	service "github.com/julianjjo/versasale-back/internal/infrastructure/service"
)

func RegisterCustomer(c *gin.Context) {
	client, _ := c.MustGet("client").(*mongo.Client)

	// Get the JSON body and decode into Customer
	var customer model.Customer
	customer.CustomerId = uuid.New().String()
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate unique email
	exists, err := service.EmailExists(client, ctx, customer.Email, "Customer")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or DocumentId already in use"})
		return
	}

	// Validate unique DocumentId
	exists, err = service.DocumentIdExists(client, ctx, customer.DocumentId, "Customer")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking DocumentId"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or DocumentId already in use"})
		return
	}

	// Hash the password
	hashedPassword, err := service.HashPassword(customer.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	customer.Password = hashedPassword

	// Save the customer to the database
	err = service.SaveCustomer(client, ctx, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer.Password = "secret"

	c.JSON(http.StatusOK, customer)
}

func RegisterSeller(c *gin.Context) {
	client, _ := c.MustGet("client").(*mongo.Client)

	// Get the JSON body and decode into Seller
	var seller model.Seller
	seller.SellerId = uuid.New().String()
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate unique email
	exists, err := service.EmailExists(client, ctx, seller.Email, "Seller")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or DocumentId already in use"})
		return
	}

	// Validate unique DocumentId
	exists, err = service.DocumentIdExists(client, ctx, seller.DocumentId, "Seller")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or DocumentId already in use"})
		return
	}

	// Hash the password
	hashedPassword, err := service.HashPassword(seller.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	seller.Password = hashedPassword

	// Save the seller to the database
	err = service.SaveSeller(client, ctx, seller)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seller)
}
