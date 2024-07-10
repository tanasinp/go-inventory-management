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
	SaveProduct(product *database.Product) error
	FindProductWithSupplierAndCategory(productID uint) (*database.Product, error)
	FindAllProductOfCategory(categoryID uint) ([]database.Product, error)
	FindAllProductOfSupplier(supplierID uint) ([]database.Product, error)
	UpdateSupplier(supplier *database.Supplier) error
}
