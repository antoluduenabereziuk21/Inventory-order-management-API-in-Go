package model

import (
	// "github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ProductName string `gorm:"not null"`
	Sku         string `gorm:"not null;unique"`
	Description string
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
}
