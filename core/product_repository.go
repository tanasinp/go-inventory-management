package core

import (
	"github.com/tanasinp/go-inventory-management/database"
)

// secondary port for product
type ProductRepository interface {
	SaveSupplier(supplier *database.Supplier) error
}
