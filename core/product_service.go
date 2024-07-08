package core

import (
	"github.com/tanasinp/go-inventory-management/database"
)

// primary port
type ProductService interface {
	CreateSupplier(supplier *database.Supplier) error
	CreateCategory(category *database.Category) error
	GetAllSupplier() ([]database.Supplier, error)
}

// business logic
type productServiceImpl struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}

func (s *productServiceImpl) CreateSupplier(supplier *database.Supplier) error {
	// business logic function
	if err := s.repo.SaveSupplier(supplier); err != nil {
		return err
	}
	return nil
}

func (s *productServiceImpl) CreateCategory(category *database.Category) error {
	// business logic function
	if err := s.repo.SaveCategory(category); err != nil {
		return err
	}
	return nil
}

func (s *productServiceImpl) GetAllSupplier() ([]database.Supplier, error) {
	suppliers, err := s.repo.FindAllSupplier()
	if err != nil {
		return nil, err
	}
	return suppliers, err
}
