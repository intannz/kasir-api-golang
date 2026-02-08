package models

type SalesReport struct {
	TotalRevenue   int            `json:"total_revenue" example:"450000"`
	TotalTransaksi int            `json:"total_transaksi" example:"15"`
	ProdukTerlaris BestSellerItem `json:"produk_terlaris"`
}

type BestSellerItem struct {
	Nama       string `json:"nama" example:"Indomie Goreng"`
	QtyTerjual int    `json:"qty_terjual" example:"50"`
}
