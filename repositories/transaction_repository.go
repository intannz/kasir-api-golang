package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api-golang/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// CreateTransaction refactor version
func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// hitung Total & Validasi Stok ---
	totalAmount := 0

	// Map untuk "Cache" data biar ga query DB 2 kali
	itemPrices := make(map[int]int)
	itemNames := make(map[int]string)

	for _, item := range items {
		var price, stock int
		var name string

		query := "SELECT name, price, stock FROM products WHERE id = $1"
		if useLock {
			query += " FOR UPDATE"
		}

		err := tx.QueryRow(query, item.ProductID).Scan(&name, &price, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("produk ID %d tidak ditemukan", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("stok '%s' kurang (sisa: %d)", name, stock)
		}

		// simpan data ke Map & hitung total
		totalAmount += price * item.Quantity
		itemPrices[item.ProductID] = price
		itemNames[item.ProductID] = name
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// update stok & insert detail
	responseDetails := make([]models.TransactionDetail, 0)

	for _, item := range items {
		price := itemPrices[item.ProductID]
		name := itemNames[item.ProductID]
		subtotal := price * item.Quantity

		// update Stok
		_, err := tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		// insert Detail Transaksi Langsung
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, item.ProductID, item.Quantity, subtotal)
		if err != nil {
			return nil, err
		}

		// susun data buat response JSON (Opsional, tapi bagus buat UX)
		responseDetails = append(responseDetails, models.TransactionDetail{
			TransactionID: transactionID,
			ProductID:     item.ProductID,
			ProductName:   name,
			Quantity:      item.Quantity,
			Subtotal:      subtotal,
		})
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     responseDetails,
	}, nil
}

func (repo *TransactionRepository) GetSalesReport(startDate, endDate string) (*models.SalesReport, error) {
	// query total pendapatan & jumlah transaksi
	querySummary := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id) 
		FROM transactions 
		WHERE created_at::date BETWEEN $1 AND $2`

	var report models.SalesReport
	err := repo.db.QueryRow(querySummary, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// query produk terlaris
	queryBestSeller := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at::date BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1`

	err = repo.db.QueryRow(queryBestSeller, startDate, endDate).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)

	// handle jika belum ada penjualan sama sekali (sql.ErrNoRows)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = models.BestSellerItem{Nama: "-", QtyTerjual: 0}
	} else if err != nil {
		return nil, err
	}

	return &report, nil
}
