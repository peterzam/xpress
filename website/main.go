package main

import (
	"Xpress/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load("../.env")
	r := routes.Handlers()

	if e != nil {
		log.Fatal("Error Loading .env file")
	}
	port := os.Getenv("SERVER_PORT")
	r.Run(":" + port)
}
