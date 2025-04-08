package models

type InvoiceItem struct {
	ID         string  `json:"id"`
	InvoiceID  string  `json:"invoice_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
