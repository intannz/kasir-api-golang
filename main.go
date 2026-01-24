package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "kasir-api-golang/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// model

type Product struct {
	ID    int    `json:"id" example:"3"`
	Name  string `json:"name" example:"Teh botol"`
	Price int    `json:"price" example:"3000"`
	Stock int    `json:"stock" example:"100"`
}

type Category struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"Soda"`
	Description string `json:"description" example:"Minuman Soda"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"Item deleted successfully"`
}

// data (in memory)

var products = []Product{
	{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
	{ID: 2, Name: "Teh Pucuk", Price: 4000, Stock: 20},
}

var categories = []Category{
	{ID: 1, Name: "Makanan ringan", Description: "Kerupuk dan cemilan"},
	{ID: 2, Name: "Minuman ringan", Description: "Es teh, es kopi, dan es buah"},
}

// helper

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func errorResponse(w http.ResponseWriter, status int, message string) {
	jsonResponse(w, status, map[string]string{"message": message})
}

// logic produk

// GetProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} Product
// @Router /api/products [get]
func getAllProducts(w http.ResponseWriter) {
	jsonResponse(w, http.StatusOK, products)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Add a new product to the inventory
// @Tags Products
// @Accept  json
// @Produce  json
// @Param body body Product true "Product Data"
// @Success 201 {object} Product "Successfully added new products!"
// @Router /api/products [post]
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProd Product
	if err := json.NewDecoder(r.Body).Decode(&newProd); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	newProd.ID = len(products) + 1
	products = append(products, newProd)
	jsonResponse(w, http.StatusCreated, newProd)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Retrieve details of a specific product
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Router /api/products/{id} [get]
func getProductByID(w http.ResponseWriter, index int) { // Hapus 'r' & 'id', cukup 'index'
	jsonResponse(w, http.StatusOK, products[index])
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update details of a specific product
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param body body Product true "Product Data"
// @Success 200 {object} Product
// @Router /api/products/{id} [put]
func updateProduct(w http.ResponseWriter, r *http.Request, index int) {
	var updateData Product
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	// update fields
	products[index].Name = updateData.Name
	products[index].Price = updateData.Price
	products[index].Stock = updateData.Stock
	jsonResponse(w, http.StatusOK, products[index])
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Remove a product from inventory
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} SuccessResponse
// @Router /api/products/{id} [delete]
func deleteProduct(w http.ResponseWriter, index int) {
	products = append(products[:index], products[index+1:]...)
	errorResponse(w, http.StatusOK, "Product deleted successfully")
}

// logic kategori

// GetCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags Categories
// @Accept  json
// @Produce  json
// @Success 200 {array} Category
// @Router /categories [get]
func getAllCategories(w http.ResponseWriter) {
	jsonResponse(w, http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Add a new category
// @Tags Categories
// @Accept  json
// @Produce  json
// @Param body body Category true "Category Data"
// @Success 201 {object} Category
// @Router /categories [post]
func createCategory(w http.ResponseWriter, r *http.Request) {
	var newCat Category
	if err := json.NewDecoder(r.Body).Decode(&newCat); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	newCat.ID = len(categories) + 1
	categories = append(categories, newCat)
	jsonResponse(w, http.StatusCreated, newCat)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update details of a specific category
// @Tags Categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param body body Category true "Category Data"
// @Success 200 {object} Category
// @Router /categories/{id} [put]
func updateCategory(w http.ResponseWriter, r *http.Request, index int) {
	var updateData Category
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}
	categories[index].Name = updateData.Name
	categories[index].Description = updateData.Description
	jsonResponse(w, http.StatusOK, categories[index])
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Remove a category
// @Tags Categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} SuccessResponse
// @Router /categories/{id} [delete]
func deleteCategory(w http.ResponseWriter, index int) {
	categories = append(categories[:index], categories[index+1:]...)
	errorResponse(w, http.StatusOK, "Category deleted successfully")
}

// routing (dispatcher)

func productsRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getAllProducts(w)
		return
	}
	if r.Method == "POST" {
		createProduct(w, r)
		return
	}
}

func productDetailRoute(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	index := -1
	for i, p := range products {
		if p.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		errorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	switch r.Method {
	case "GET":
		getProductByID(w, index)
	case "PUT":
		updateProduct(w, r, index)
	case "DELETE":
		deleteProduct(w, index)
	default:
		errorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func categoriesRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getAllCategories(w)
		return
	}
	if r.Method == "POST" {
		createCategory(w, r)
		return
	}
}

func categoryDetailRoute(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	index := -1
	for i, c := range categories {
		if c.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		errorResponse(w, http.StatusNotFound, "Category not found")
		return
	}

	switch r.Method {
	case "GET":
		jsonResponse(w, http.StatusOK, categories[index])
	case "PUT":
		updateCategory(w, r, index)
	case "DELETE":
		deleteCategory(w, index)
	default:
		errorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// @title Kasir API Intan
// @version 1.0
// @description API Kasir Toko (In-Memory)
// @contact.name Intan Maharani
// @contact.email intan.maharani6763@gmail.com
// @BasePath /
func main() {
	// route produk
	http.HandleFunc("/api/products", productsRoute)
	http.HandleFunc("/api/products/", productDetailRoute)

	// route kategori
	http.HandleFunc("/categories", categoriesRoute)
	http.HandleFunc("/categories/", categoryDetailRoute)

	// swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port " + port)
	http.ListenAndServe(":"+port, nil)
}
