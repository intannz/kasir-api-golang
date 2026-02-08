package main

import (
	"fmt"
	"log"
	"net/http"

	"kasir-api-golang/database"
	"kasir-api-golang/handlers"
	"kasir-api-golang/repositories"
	"kasir-api-golang/services"
	"kasir-api-golang/util"

	_ "kasir-api-golang/docs"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Kasir API Intan
// @version 1.0
// @description API Kasir Toko dengan Database (Supabase)
// @contact.name Intan Maharani
// @contact.email intan.maharani6763@gmail.com
// @BasePath /
func main() {
	// load Config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// konek Database
	db, err := database.InitDB(config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer db.Close() // tutup koneksi saat aplikasi berhenti

	// Setup Layers
	// Setup product Layer
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Setup category Layer
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Setup routes
	// Route produk
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	// Route kategori
	http.HandleFunc("/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryByID)

	// Route checkout & report
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout)
	http.HandleFunc("/api/report", transactionHandler.HandleReport)

	// swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Start Server
	fmt.Println("Server running on address:", config.ServerAddress)
	log.Fatal(http.ListenAndServe(config.ServerAddress, nil))
}
