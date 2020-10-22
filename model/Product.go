package model

import (
	"time"
)

type Product struct {
	ID        uint64    `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null"`
	Name      string    `gorm:"not null;unique"`
	Price     float64   `gorm:"type:decimal(10,2);not null;unique;default:0.0"`
	Quantity  uint16    `gorm:"default:0;unsigned"`
	State     string    `gorm:"not null"`
}
