package router

import (
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(productHandler *handler.ProductHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/products/{id}", productHandler.GetProductByID)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Get("/products", productHandler.GetAllProducts)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	return r
}
