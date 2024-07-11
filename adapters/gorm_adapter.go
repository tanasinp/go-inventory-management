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

func (r *gormProductRepository) FindProductByID(productID uint) (*database.Product, error) {
	var product database.Product
	if result := r.db.Preload("Supplier").Preload("Categories").First(&product, productID); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *gormProductRepository) FindAllProductOfCategory(categoryID uint) ([]database.Product, error) {
	var products []database.Product
	result := r.db.Preload("Supplier").Preload("Categories").Joins("JOIN product_categories on product_categories.product_id = products.id").
		Where("product_categories.category_id = ?", categoryID).
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *gormProductRepository) FindAllProductOfSupplier(supplierID uint) ([]database.Product, error) {
	var products []database.Product
	result := r.db.Preload("Supplier").Preload("Categories").Joins("Join suppliers on suppliers.id = products.supplier_id").
		Where("suppliers.id = ?", supplierID).
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *gormProductRepository) UpdateSupplier(supplier *database.Supplier) error {
	result := r.db.Model(&supplier).Updates(supplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *gormProductRepository) FindAllProduct() ([]database.Product, error) {
	var products []database.Product
	if result := r.db.Preload("Supplier").Preload("Categories").Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *gormProductRepository) UpdateProductByID(productID uint, updatedProduct *database.Product) error {
	var product database.Product
	if err := r.db.Preload("Categories").Preload("Supplier").First(&product, productID).Error; err != nil {
		return err
	}

	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.SupplierID = updatedProduct.SupplierID
	// Clear existing categories
	if err := r.db.Model(&product).Association("Categories").Clear(); err != nil {
		return err
	}
	if len(updatedProduct.Categories) > 0 {
		var categories []database.Category
		for _, cat := range updatedProduct.Categories {
			var category database.Category
			if err := r.db.First(&category, cat.ID).Error; err != nil {
				return err
			}
			categories = append(categories, category)
		}
		product.Categories = categories
	}

	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&product).Error
}

func (r *gormProductRepository) DeleteProductByID(productID uint) error {
	var product database.Product
	if err := r.db.Preload("Categories").Preload("Supplier").First(&product, productID).Error; err != nil {
		return err
	}
	if err := r.db.Model(&product).Association("Categories").Clear(); err != nil {
		return err
	}
	if err := r.db.Unscoped().Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
