package main

import (
	"fmt"
	"xpress/routes"
	"xpress/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(utils.GetEnv("GIN_MODE"))
	r := routes.Handlers()
	err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}
	port := utils.GetEnv("WEB_PORT")
	fmt.Printf("The web server is running at %s:%s\n", utils.GetEnv("WEB_ADDRESS"), utils.GetEnv("WEB_PORT"))
	err = r.Run(":" + port)
	if err != nil {
		panic(err)
	}

}
