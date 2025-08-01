package main

import (
	"fmt"
	"go-api/api/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// ENV Initialization
	err := godotenv.Load()
	if err != nil {
		panic("Error when initializing environment variables: " + err.Error())
	} else {
		fmt.Println("ENV variables loaded!")
	}

	// API Initialization
	if os.Getenv("MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	server := gin.Default()

	// Routes Configuration
	route.ConfigureHealthRoutes(server)

	err = server.Run(":" + os.Getenv("API_PORT"))

	if err != nil {
		panic("Error when starting server: " + err.Error())
	}
}
