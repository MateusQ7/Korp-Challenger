package routes

import (
	"billing-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/invoices", handlers.CreateInvoice)
	r.PUT("/invoices/:id/print", handlers.PrintInvoiceHandler)
	r.GET("/invoices/:id", handlers.GetInvoiceByID)
	r.GET("/invoices", handlers.GetAllInvoices)
}
