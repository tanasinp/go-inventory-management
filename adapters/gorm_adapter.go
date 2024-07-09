package adapters

import (
	"github.com/tanasinp/go-inventory-management/core"
	"github.com/tanasinp/go-inventory-management/database"
	"gorm.io/gorm"
)

// Secondary adapter
type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) core.ProductRepository {
	return &gormProductRepository{db: db}
}

func (r *gormProductRepository) SaveSupplier(supplier *database.Supplier) error {
	if result := r.db.Create(supplier); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *gormProductRepository) SaveCategory(category *database.Category) error {
	if result := r.db.Create(category); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *gormProductRepository) FindAllSupplier() ([]database.Supplier, error) {
	var suppliers []database.Supplier
	if result := r.db.Find(&suppliers); result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

func (r *gormProductRepository) FindAllCategory() ([]database.Category, error) {
	var categories []database.Category
	if result := r.db.Find(&categories); result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *gormProductRepository) SaveProduct(product *database.Product) error {
	if result := r.db.Create(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *gormProductRepository) FindProductWithSupplier(productID uint) (*database.Product, error) {
	var product database.Product
	if result := r.db.Preload("Supplier").First(&product, productID); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
