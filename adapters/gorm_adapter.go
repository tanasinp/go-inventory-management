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

func (r *gormProductRepository) FindProductWithSupplierAndCategory(productID uint) (*database.Product, error) {
	var product database.Product
	if result := r.db.Preload("Supplier").Preload("Categories").First(&product, productID); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *gormProductRepository) FindAllProductOfCategory(categoryID uint) ([]database.Product, error) {
	var products []database.Product
	result := r.db.Joins("JOIN product_categories on product_categories.product_id = products.id").
		Where("product_categories.category_id = ?", categoryID).
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
