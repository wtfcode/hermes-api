package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wtfcode/hermes-api/pkg/database"
	"github.com/wtfcode/hermes-api/pkg/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connect()
	database.Migrate(db)

	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "7000"
	}

	app := server.NewServer(db)
	app.Listen(":" + PORT)
}
