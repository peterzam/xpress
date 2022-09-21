package main

import (
	"Xpress/routes"
	"Xpress/utils"
)

func main() {

	r := routes.Handlers()
	err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}
	port := utils.GetEnv("WEB_PORT")
	r.Run(":" + port)
}
