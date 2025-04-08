package repository

import (
	"fmt"
	"log"
	"stock-service/database"
	"stock-service/models"
)

func GetProducts() ([]models.Product, error) {
	rows, err := database.DB.Query(`
		SELECT id, name, price, stock, created_at, updated_at
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Stock,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	var p models.Product
	err := database.DB.QueryRow(`
		SELECT id, name, price, stock, created_at, updated_at
		FROM products
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.Stock,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func CreateProduct(product *models.Product) error {
	query := `
		INSERT INTO products (name, price, stock)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := database.DB.QueryRow(
		query,
		product.Name,
		product.Price,
		product.Stock,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	return err
}

func DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`

	_, err := database.DB.Exec(query, id)
	return err
}

func DecreaseStock(productID, quantity int) error {
	query := `
		UPDATE products 
		SET stock = stock - $1
		WHERE id = $2 AND stock >= $1
	`

	result, err := database.DB.Exec(query, quantity, productID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("insufficient stock for product %d", productID)
	}

	return nil
}

func UpdateStock(id int, quantityToSubtract int) error {
	query := `
		UPDATE products
		SET stock = stock - $1, updated_at = NOW()
		WHERE id = $2 AND stock >= $1
	`

	result, err := database.DB.Exec(query, quantityToSubtract, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("Product %d - trying to subtract %d units - rows affected: %d\n", id, quantityToSubtract, rowsAffected)
	if rowsAffected == 0 {
		return fmt.Errorf("insufficient stock for product %d", id)
	}

	return nil
}
