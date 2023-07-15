package main

import (
	"day_eight/config"
	"day_eight/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
