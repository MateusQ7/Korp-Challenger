package routes

import (
	"stock-service/handlers"

	"github.com/gin-gonic/gin"
)

func ProductsRoutes(router *gin.Engine) {
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProductByID)
	router.POST("/products", handlers.CreateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)
	router.PUT("/products/:id/decrease", handlers.DecreaseStock)
	router.PUT("/products/:id/stock", handlers.UpdateStock)
}
