package services

import (
	"kasir-api-golang/models"
	"kasir-api-golang/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GET ALL
func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

// CREATE
func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

// GET BY ID
func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// UPDATE
func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

// DELETE
func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
