package handlers

import (
	"encoding/json"
	"kasir-api-golang/models"
	"kasir-api-golang/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// DISPATCHERS (Jembatan / Satpam)

// dispatcher route "/api/products" (ngga pakai ID)
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// dispatcher untuk Route "/api/products/" (pakai ID)
func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	// ambil ID dari URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r, id)
	case http.MethodPut:
		h.Update(w, r, id)
	case http.MethodDelete:
		h.Delete(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// LOGIC HANDLERS

// GetAll godoc
// @Summary      Ambil semua produk
// @Description  Mengambil daftar semua produk
// @Tags         2. products
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Product
// @Router       /api/products [get]
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Create godoc
// @Summary      Tambah produk baru
// @Description  Menambahkan data produk ke database
// @Tags         2. products
// @Accept       json
// @Produce      json
// @Param        product body models.Product true "Data Produk"
// @Success      201  {object}  models.Product
// @Router       /api/products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	// decode JSON Body
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// panggil Service
	if err := h.service.Create(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetByID godoc
// @Summary      Ambil produk berdasarkan ID
// @Description  Mendapatkan detail satu produk
// @Tags         2. products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.Product
// @Router       /api/products/{id} [get]
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	product, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Update godoc
// @Summary      Update data produk
// @Description  Mengubah data produk yang sudah ada
// @Tags         2. products
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Product ID"
// @Param        product  body      models.Product  true  "Data Update"
// @Success      200      {object}  models.Product
// @Router       /api/products/{id} [put]
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// set ID dari URL ke struct
	product.ID = id

	if err := h.service.Update(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Delete godoc
// @Summary      Hapus produk
// @Description  Menghapus produk berdasarkan ID
// @Tags         2. products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]string
// @Router       /api/products/{id} [delete]
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}
