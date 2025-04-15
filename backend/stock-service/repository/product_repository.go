package repository

import (
	"fmt"
	"stock-service/database"
	"stock-service/models"

	"gorm.io/gorm"
)

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Find(&products)
	return products, result.Error
}

func GetProductByID(id int) (models.Product, error) {
	var product models.Product
	result := database.DB.First(&product, id)
	return product, result.Error
}

func CreateProduct(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

func DeleteProduct(id int) error {
	result := database.DB.Delete(&models.Product{}, id)
	return result.Error
}

func DecreaseStock(productID, quantity int) error {
	result := database.DB.Model(&models.Product{}).
		Where("id = ? AND stock >= ?", productID, quantity).
		UpdateColumn("stock", gorm.Expr("stock - ?", quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("insufficient stock for product %d", productID)
	}

	return nil
}

func UpdateStock(product *models.Product) error {
	result := database.DB.Save(product)
	return result.Error
}
