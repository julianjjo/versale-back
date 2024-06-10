package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Container contains environment variables for the application, database, cache, token, and http server
type (
	Container struct {
		MongoDB *MongoDB
		HTTP    *HTTP
	}
	// Database contains all the environment variables for the database
	MongoDB struct {
		Connection string
	}
	// HTTP contains all the environment variables for the http server
	HTTP struct {
		URL    string
		Port   string
		Secret string
	}
)

// New creates a new container instance
func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	mongodb := &MongoDB{
		Connection: os.Getenv("MONGODB_CONNECTION_STRING"),
	}

	http := &HTTP{
		URL:    os.Getenv("HTTP_URL"),
		Port:   os.Getenv("HTTP_PORT"),
		Secret: os.Getenv("SECRET"),
	}

	return &Container{
		mongodb,
		http,
	}, nil
}
