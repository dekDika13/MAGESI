package main

import (
	"log"
	"magesi/database"
	"magesi/router"
	"os"

	"github.com/labstack/echo"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	e := echo.New()

	router.New(e, database.DB)

	port := os.Getenv("PORT")

	_ = e.Start(port)
}
