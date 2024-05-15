package controllers

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	config "github.com/julianjjo/versasale-back/internal/infrastructure/config"
	service "github.com/julianjjo/versasale-back/internal/infrastructure/service"
)

func LoginCustomer(c *gin.Context) {
	client, _ := c.MustGet("client").(*mongo.Client)

	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var responseLogin struct {
		Menssage string `json:"message"`
		Token    string `json:"token"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var customer service.Customer // Assuming the user can be a customer, adapt if needed
	errMongoDB := client.Database("versasale").Collection("customer").FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&customer)
	if errMongoDB != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !service.CheckPasswordHash(loginData.Password, customer.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": customer.CustomerId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.HTTP.Secret))

	responseLogin.Menssage = "Login successful"
	responseLogin.Token = tokenString

	c.JSON(http.StatusOK, responseLogin)
}

func LoginSeller(c *gin.Context) {
	client, _ := c.MustGet("client").(*mongo.Client)
	ctx, _ := c.MustGet("ctx").(context.Context)

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var seller service.Seller // Assuming the user can be a seller, adapt if needed
	err := client.Database("versasale").Collection("customer").FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&seller)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !service.CheckPasswordHash(loginData.Password, seller.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
