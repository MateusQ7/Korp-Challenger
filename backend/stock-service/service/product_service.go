package service

import (
	"errors"
	"fmt"
	"log"
	"stock-service/models"
	"stock-service/repository"
	"strings"
)

func CreateProduct(product *models.Product) error {

	product.Name = strings.TrimSpace(product.Name)

	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}

	if product.Price < 0 {
		return errors.New("product price cannot be negative")
	}

	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	log.Printf("Creating product: %+v\n", product)

	return repository.CreateProduct(product)
}

func DecreaseProductStock(productID, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	product, err := repository.GetProductByID(productID)
	if err != nil {
		return fmt.Errorf("failed to find product: %w", err)
	}

	if product.Stock < quantity {
		return fmt.Errorf("insufficient stock for product %d", productID)
	}

	product.Stock -= quantity

	err = repository.UpdateStock(&product)
	if err != nil {
		return fmt.Errorf("failed to update stock: %w", err)
	}

	return nil
}
