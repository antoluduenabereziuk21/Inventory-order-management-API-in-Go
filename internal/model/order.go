package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID   uint
	Customer     Customer      `gorm:"foreignKey:CustomerID"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
	Status       OrderStatus   `gorm:"not null"`
}

type OrderStatus int

const (
	StatusPending OrderStatus = iota
	StatusProcessing
	StatusShipped
	StatusDelivered
	StatusCancelled
)
