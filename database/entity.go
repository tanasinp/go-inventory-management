package database

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SupplierID  int
	Supplier    Supplier
	Categories  []Category `gorm:"many2many:product_categories;"`
}

type Supplier struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Contact string
}

type Category struct {
	gorm.Model
	Name     string    `gorm:"unique"`
	Products []Product `gorm:"many2many:product_categories;"`
}

type ProductCategory struct {
	ProductID  int
	Product    Product
	CategoryID int
	Category   Category
}
