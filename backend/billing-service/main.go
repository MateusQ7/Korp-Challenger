package main

import (
	"billing-service/database"
	"billing-service/middlewares"
	"billing-service/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	database.RunMigrations(db)

	r := gin.Default()

	middlewares.SetupCors(r)
	routes.SetupRoutes(r)

	fmt.Println("Billing server running on port 8082...")
	r.Run(":8082")
}
