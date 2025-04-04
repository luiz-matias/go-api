package main

import (
	"fmt"
	"go-api/api/controller"
	"go-api/api/route"
	"go-api/domain/use_case"
	"go-api/infra/db"
	"go-api/infra/repository"
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

	// Database Initialization
	dbConnection := db.GetConnection()

	// Dependency Injection
	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := use_case.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	// Routes Configuration
	route.ConfigureHealthRoutes(server)
	route.ConfigureProductRoutes(server, &productController)

	err = server.Run(":" + os.Getenv("API_PORT"))

	if err != nil {
		panic("Error when starting server: " + err.Error())
	}
}
