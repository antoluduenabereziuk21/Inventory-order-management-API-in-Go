package repository

import (
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/model"
)

type ProductRepository interface {
	// GetProductByID retrieves a product by its ID.
	GetProductByID(id uint) (*model.Product, error)
	GetAllProducts() ([]*model.Product, error)
	// CreateProduct creates a new product.
	CreateProduct(product *model.Product) error
	// UpdateProduct updates an existing product.
	UpdateProduct(product *model.Product) error
	// DeleteProduct deletes a product by its ID.
	DeleteProduct(id uint) error
}
