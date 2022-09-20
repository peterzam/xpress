package main

import (
	"Xpress/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	r := routes.Handlers()

	if e != nil {
		log.Fatal("Error Loading .env file")
	}
	// port := os.Getenv("PORT")
	r.Run(":" + "8080")
}
