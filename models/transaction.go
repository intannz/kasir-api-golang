package models

import "time"

// representasi struk belanjaan
type Transaction struct {
	ID          int                 `json:"id" example:"101"`
	TotalAmount int                 `json:"total_amount" example:"45000"`
	CreatedAt   time.Time           `json:"created_at" example:"2026-02-01T14:00:00Z"`
	Details     []TransactionDetail `json:"details"`
}

// rincian barang apa saja yang dibeli dalam 1 struk
type TransactionDetail struct {
	ID            int    `json:"id" example:"505"`
	TransactionID int    `json:"transaction_id" example:"101"`
	ProductID     int    `json:"product_id" example:"2"`
	ProductName   string `json:"product_name,omitempty" example:"Indomie Goreng"`
	Quantity      int    `json:"quantity" example:"5"`
	Subtotal      int    `json:"subtotal" example:"15000"`
}

// inputan simpel dari Kasir (hanya butuh ID dan Jumlah)
type CheckoutItem struct {
	ProductID int `json:"product_id" example:"1"`
	Quantity  int `json:"quantity" example:"2"`
}

// pembungkus (wrapper) biar JSON-nya rapi ada key "items"
type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}
