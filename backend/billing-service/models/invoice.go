package models

import "time"

type Invoice struct {
	ID        string        `json:"id"`
	Number    string        `json:"number"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	Products  []InvoiceItem `json:"products"`
}
