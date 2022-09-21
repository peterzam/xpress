package main

import (
	"Xpress/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	r := routes.Handlers()

	if godotenv.Load("../.env") != nil {
		panic("Error Loading .env file")
	}
	port := os.Getenv("WEB_PORT")
	r.Run(":" + port)
}
