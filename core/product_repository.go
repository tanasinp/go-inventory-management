package core

import (
	"github.com/tanasinp/go-inventory-management/database"
)

// secondary port for product
type ProductRepository interface {
	SaveSupplier(supplier *database.Supplier) error
	SaveCategory(category *database.Category) error
	FindAllSupplier() ([]database.Supplier, error)
	FindAllCategory() ([]database.Category, error)
}
