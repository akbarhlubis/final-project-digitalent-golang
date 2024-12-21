package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// App
	AppEnv  string
	AppURL  string
	AppHost string
	AppPort string

	// Database
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string

	// Frontend
	FrontendURL  string
	FrontendHost string
	FrontendPort string
	FrontendFull string
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	AppEnv = os.Getenv("APP_ENV")
	AppURL = os.Getenv("APP_URL")
	AppHost = os.Getenv("APP_HOST")
	AppPort = os.Getenv("APP_PORT")

	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_DATABASE")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")

	FrontendURL = os.Getenv("FRONTEND_URL")
	FrontendHost = os.Getenv("FRONTEND_HOST")
	FrontendPort = os.Getenv("FRONTEN_PORT")
	FrontendFull = os.Getenv("FRONTEND_FULL_HOST_PORT")
}
