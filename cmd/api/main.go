package main

import (
	"log"
	"wowza/internal/app"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment variables")
	}
}

// @title Wowza API
// @version 1.0
// @description This is a sample server for a Wowza API.
// @host localhost:8080
// @BasePath /
func main() {
	app.Run()
}
