package service

import (
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/model"
)

type ProductService interface {
	GetProductByID(id uint) (*model.Product, error)
	GetAllProducts() ([]*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(id uint, product *model.Product) error
	DeleteProduct(id uint) error
}
