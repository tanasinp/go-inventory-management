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
