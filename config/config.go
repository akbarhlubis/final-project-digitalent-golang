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
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
}
