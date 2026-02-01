package models

type Product struct {
	ID           int    `json:"id" example:"1"`
	Name         string `json:"name" example:"Nasi Goreng Spesial"`
	Price        int    `json:"price" example:"25000"`
	Stock        int    `json:"stock" example:"100"`
	CategoryID   int    `json:"categoryId" example:"1"`
	CategoryName string `json:"categoryName,omitempty" example:"Makanan Berat"`
}
