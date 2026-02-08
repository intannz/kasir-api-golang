package services

import (
	"kasir-api-golang/models"
	"kasir-api-golang/repositories"
	"time"
)

// definisi struct
type TransactionService struct {
	repo *repositories.TransactionRepository
}

// constructor
func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// method Checkout
func (s *TransactionService) Checkout(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items, useLock)
}

func (s *TransactionService) GetReport(start, end string) (*models.SalesReport, error) {
	// Jika parameter kosong, default ke HARI INI
	if start == "" || end == "" {
		now := time.Now().Format("2006-01-02")
		start = now
		end = now
	}
	return s.repo.GetSalesReport(start, end)
}
