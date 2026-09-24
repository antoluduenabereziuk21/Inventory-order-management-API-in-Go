package handler

import (
	"net/http"

	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/service"

	"strconv"

	"github.com/go-chi/chi/v5"

	"encoding/json"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	IdUint, err := uint64(strconv.ParseUint(id, 10, 64))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product, err := h.service.GetProductByID(IdUint)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
