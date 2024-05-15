package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	routers "github.com/julianjjo/versasale-back/internal/adapter/router"
	config "github.com/julianjjo/versasale-back/internal/infrastructure/config"
	repository "github.com/julianjjo/versasale-back/internal/infrastructure/repository"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Connect to MongoDB
	client, ctx, cancel, err := repository.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Successfully connected and pinged MongoDB.")

	router := routers.SetupRouter(client, ctx)
	router.Use(gin.Recovery())

	router.GET("/panic", func(c *gin.Context) {
		panic("Something went wrong!")
	})

	router.Run(config.HTTP.URL + ":" + config.HTTP.Port)
}
