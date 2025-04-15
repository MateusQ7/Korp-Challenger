package main

import (
	"fmt"
	"stock-service/database"
	"stock-service/middlewares"
	"stock-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	database.RunMigrations(database.RawDB())

	r := gin.Default()

	middlewares.SetupCors(r)

	routes.ProductsRoutes(r)

	fmt.Println("Server is running on port 8081")
	r.Run(":8081")
}
