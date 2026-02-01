package services

import (
	"kasir-api-golang/models"
	"kasir-api-golang/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GET ALL
func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

// CREATE
func (s *CategoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

// GET BY ID
func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

// UPDATE
func (s *CategoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}

// DELETE
func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
