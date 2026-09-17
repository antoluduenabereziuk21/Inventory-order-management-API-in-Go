package model

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int
	Price     float64
}
