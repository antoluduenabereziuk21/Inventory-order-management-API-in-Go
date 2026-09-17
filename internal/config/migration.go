package config

import (
	"log"

	"github.com/antoluduenabereziuk21/Inventory-order-management-API-in-Go/internal/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {

	entities := []interface{}{
		&model.Customer{},
		&model.Product{},
		&model.Order{},
		&model.OrderDetail{},
	}
	for _, entity := range entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			log.Fatalf("Failed to migrate entity %T: %v", entity, err)
		}
	}
}
