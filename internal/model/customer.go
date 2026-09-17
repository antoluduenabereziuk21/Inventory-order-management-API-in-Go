package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null;unique"`
	Address string
	Phone   string
	Orders  []Order `gorm:"foreignKey:CustomerID"`
}
