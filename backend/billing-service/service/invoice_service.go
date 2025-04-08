package service

import (
	"billing-service/dto"
	"billing-service/models"
	"billing-service/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const stockServiceBaseURL = "http://localhost:8081/products/"

func validateStock(productID int, requiredQty int) (*dto.ProductResponse, error) {
	resp, err := http.Get(stockServiceBaseURL + strconv.Itoa(productID))
	if err != nil {
		return nil, fmt.Errorf("failed to communicate with stock service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product %d not found (status: %d)", productID, resp.StatusCode)
	}

	var product dto.ProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, fmt.Errorf("error decoding stock service response: %v", err)
	}

	if product.Stock < requiredQty {
		return nil, fmt.Errorf("insufficient stock for product %d (%s): required %d, available %d",
			product.ID, product.Name, requiredQty, product.Stock)
	}

	return &product, nil
}

func CreateInvoice(invoice models.Invoice) (string, error) {
	for i := range invoice.Products {
		product, err := validateStock(invoice.Products[i].ProductID, invoice.Products[i].Quantity)
		if err != nil {
			return "", err
		}

		invoice.Products[i].UnitPrice = product.Price
		invoice.Products[i].TotalPrice = float64(invoice.Products[i].Quantity) * product.Price
	}

	invoiceID, err := repository.SaveInvoice(invoice)
	if err != nil {
		return "", fmt.Errorf("failed to save invoice: %v", err)
	}

	return invoiceID, nil
}

func PrintInvoice(invoiceID int) error {
	invoice, err := repository.GetInvoiceByID(invoiceID)
	if err != nil {
		return fmt.Errorf("invoice not found: %v", err)
	}

	if invoice.Status == "fechada" {
		return fmt.Errorf("invoice is already closed")
	}

	for _, item := range invoice.Products {
		_, err := validateStock(item.ProductID, item.Quantity)
		if err != nil {
			return err
		}

		payload := map[string]int{"quantity": item.Quantity}
		jsonBody, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal stock update payload: %v", err)
		}

		url := stockServiceBaseURL + strconv.Itoa(item.ProductID) + "/stock"
		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return fmt.Errorf("error building stock update request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to update stock for product %d: %v", item.ProductID, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("stock service returned status %d for product %d", resp.StatusCode, item.ProductID)
		}
	}

	err = repository.UpdateInvoiceStatus(invoiceID, "fechada")
	if err != nil {
		return fmt.Errorf("failed to close invoice: %v", err)
	}

	return nil
}

func EnsureInvoiceIsOpen(invoiceID int) error {
	invoice, err := repository.GetInvoiceByID(invoiceID)
	if err != nil {
		return fmt.Errorf("invoice not found")
	}

	if invoice.Status == "fechada" {
		return fmt.Errorf("action not allowed: invoice is closed")
	}

	return nil
}
