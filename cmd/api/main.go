package main

import (
	"net/http"

	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/config"
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/handler"
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/router"
	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/service"
)

func main() {
	productService := service.NewProductService(config.DBConn)
	productHandler := handler.NewProductHandler(productService)
	r := router.NewRouter(productHandler)

	http.ListenAndServe(":8080", r)
}
