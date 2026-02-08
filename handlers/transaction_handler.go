package handlers

import (
	"encoding/json"
	"net/http"

	"kasir-api-golang/models"
	"kasir-api-golang/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Checkout godoc
// @Summary      Proses Transaksi (Checkout)
// @Description  User membeli barang, stok berkurang, transaksi tercatat
// @Tags         3. Transactions
// @Accept       json
// @Produce      json
// @Param        request body models.CheckoutRequest true "Item yang dibeli"
// @Success      201  {object}  models.Transaction
// @Router       /api/checkout [post]
func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest

	// decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	useLock := true

	// panggil service dengan parameter useLock
	transaction, err := h.service.Checkout(req.Items, useLock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// sukses! kembalikan data transaksi
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// GetReport godoc
// @Summary      Ambil Laporan Penjualan
// @Description  Mendapatkan total omzet, jumlah transaksi, dan produk terlaris. Default: Hari ini.
// @Tags         3. Transactions
// @Accept       json
// @Produce      json
// @Param        start_date query string false "Tanggal Mulai (YYYY-MM-DD)"
// @Param        end_date   query string false "Tanggal Akhir (YYYY-MM-DD)"
// @Success      200 {object} models.SalesReport
// @Router       /api/report [get]
func (h *TransactionHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	report, err := h.service.GetReport(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
