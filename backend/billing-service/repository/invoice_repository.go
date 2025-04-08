package repository

import (
	"billing-service/database"
	"billing-service/models"
	"fmt"
)

func SaveInvoice(invoice models.Invoice) (string, error) {
	invoiceTx, err := database.DB.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to start transaction: %v", err)
	}

	rollback := true

	defer func() {
		if rollback {
			invoiceTx.Rollback()
		}
	}()

	var invoiceID string
	if err := invoiceTx.QueryRow(`
		INSERT INTO invoices (number, status) 
		VALUES ($1, $2) RETURNING id
	`, invoice.Number, invoice.Status).Scan(&invoiceID); err != nil {
		return "", fmt.Errorf("failed to save invoice: %v", err)
	}

	for _, p := range invoice.Products {
		p.TotalPrice = float64(p.Quantity) * p.UnitPrice

		if _, err := invoiceTx.Exec(`
			INSERT INTO invoice_items (invoice_id, product_id, quantity, unit_price, total_price) 
			VALUES ($1, $2, $3, $4, $5)
		`, invoiceID, p.ProductID, p.Quantity, p.UnitPrice, p.TotalPrice); err != nil {
			return "", fmt.Errorf("failed to associate product %d with invoice: %v", p.ProductID, err)
		}
	}

	if err := invoiceTx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	rollback = false
	return invoiceID, nil
}

func UpdateInvoiceStatus(invoiceID int, status string) error {
	query := `UPDATE invoices SET status = $1 WHERE id = $2`
	_, err := database.DB.Exec(query, status, invoiceID)
	return err
}

func GetInvoiceByID(id int) (models.Invoice, error) {
	var invoice models.Invoice

	query := `
		SELECT id, number, status, created_at
		FROM invoices
		WHERE id = $1
	`
	err := database.DB.QueryRow(query, id).Scan(
		&invoice.ID,
		&invoice.Number,
		&invoice.Status,
		&invoice.CreatedAt,
	)
	if err != nil {
		return invoice, fmt.Errorf("failed to retrieve invoice: %w", err)
	}

	itemsQuery := `
		SELECT id, invoice_id, product_id, quantity, unit_price, total_price
		FROM invoice_items
		WHERE invoice_id = $1
	`

	rows, err := database.DB.Query(itemsQuery, id)
	if err != nil {
		return invoice, fmt.Errorf("failed to retrieve invoice items: %w", err)
	}
	defer rows.Close()

	var items []models.InvoiceItem

	for rows.Next() {
		var item models.InvoiceItem
		err := rows.Scan(
			&item.ID,
			&item.InvoiceID,
			&item.ProductID,
			&item.Quantity,
			&item.UnitPrice,
			&item.TotalPrice,
		)
		if err != nil {
			return invoice, fmt.Errorf("failed to read invoice item: %w", err)
		}
		items = append(items, item)
	}

	invoice.Products = items

	return invoice, nil
}

func GetAllInvoices() ([]models.Invoice, error) {
	query := `
		SELECT id, number, status, created_at
		FROM invoices
		ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch invoices: %w", err)
	}
	defer rows.Close()

	var invoices []models.Invoice

	for rows.Next() {
		var invoice models.Invoice
		err := rows.Scan(
			&invoice.ID,
			&invoice.Number,
			&invoice.Status,
			&invoice.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan invoice: %w", err)
		}

		itemsQuery := `
			SELECT id, invoice_id, product_id, quantity, unit_price, total_price
			FROM invoice_items
			WHERE invoice_id = $1
		`
		itemRows, err := database.DB.Query(itemsQuery, invoice.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch invoice items for invoice %s: %w", invoice.ID, err)
		}

		var items []models.InvoiceItem
		for itemRows.Next() {
			var item models.InvoiceItem
			err := itemRows.Scan(
				&item.ID,
				&item.InvoiceID,
				&item.ProductID,
				&item.Quantity,
				&item.UnitPrice,
				&item.TotalPrice,
			)
			if err != nil {
				itemRows.Close()
				return nil, fmt.Errorf("failed to scan invoice item: %w", err)
			}
			items = append(items, item)
		}
		itemRows.Close()

		invoice.Products = items
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}
