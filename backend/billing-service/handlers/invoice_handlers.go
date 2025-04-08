package handlers

import (
	"billing-service/models"
	"billing-service/repository"
	"billing-service/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	var newInvoice models.Invoice

	if err := c.ShouldBindJSON(&newInvoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	newInvoice.Status = "aberta"
	now := time.Now()
	newInvoice.CreatedAt = now

	invoiceID, err := service.CreateInvoice(newInvoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create invoice",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Invoice created successfully",
		"id":      invoiceID,
	})
}

func PrintInvoiceHandler(c *gin.Context) {
	idParam := c.Param("id")
	invoiceID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = service.PrintInvoice(invoiceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice printed successfully!"})
}

func GetInvoiceByID(c *gin.Context) {
	idParam := c.Param("id")
	invoiceID, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	invoice, err := repository.GetInvoiceByID(invoiceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Invoice not found",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func GetAllInvoices(c *gin.Context) {
	invoices, err := repository.GetAllInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve invoices",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, invoices)
}
