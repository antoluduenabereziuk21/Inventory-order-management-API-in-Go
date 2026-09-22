package repository

import (
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/model"
	"gorm.io/gorm"
)

type ProductRepositoryImp struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImp{db: db}
}
func (r *ProductRepositoryImp) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImp) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImp) CreateProduct(product *model.Product) error {
	err := r.db.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImp) UpdateProduct(product *model.Product) error {
	err := r.db.Updates(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImp) DeleteProduct(id uint) error {
	err := r.db.Delete(&model.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
