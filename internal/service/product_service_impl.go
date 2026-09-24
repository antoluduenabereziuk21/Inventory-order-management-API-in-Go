package service

import (
	"errors"
	"fmt"

	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/model"
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/repository"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	productRepo repository.ProductRepository
	db          *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		productRepo: repository.NewProductRepository(db),
		db:          db,
	}
}

func (s *ProductServiceImpl) GetProductByID(id uint) (*model.Product, error) {
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) GetAllProducts() ([]*model.Product, error) {
	products, err := s.productRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductServiceImpl) CreateProduct(product *model.Product) error {
	if product.ProductName == "" || product.Sku == "" || product.Price <= 0 || product.Stock < 0 {
		fmt.Println("Invalid product data:", product)
		return errors.New("invalid product data")
	}
	err := s.productRepo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceImpl) UpdateProduct(id uint, product *model.Product) error {
	productFromDB, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return err
	}
	if product.ProductName == "" || product.Sku == "" || product.Price <= 0 || product.Stock < 0 {
		fmt.Println("Invalid product data:", product)
		return errors.New("invalid product data")
	}
	if productFromDB.Sku != product.Sku {
		return errors.New("SKU cannot be changed")
	}

	errUpdate := s.productRepo.UpdateProduct(id, product)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (s *ProductServiceImpl) DeleteProduct(id uint) error {
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		fmt.Println("Product not found with ID:", id)
		return errors.New("product not found")
	}
	if product.Stock > 0 {
		fmt.Println("Cannot delete product with stock greater than zero. Product ID:", id)
		return errors.New("cannot delete product with stock greater than zero")
	}
	errDel := s.productRepo.DeleteProduct(id)
	if errDel != nil {
		return errDel
	}
	return nil
}
