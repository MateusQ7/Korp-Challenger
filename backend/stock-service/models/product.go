package models

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type=varchar(255);not null" json:"name"`
	Price     float64   `gorm:"not null" json:"price"`
	Stock     int       `gorm:"not null" json:"stock"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
