package main

import (
	"encoding/json" //mengubah data Go ke JSON dan sebaliknya
	"fmt"           //untuk print ke terminal
	"net/http"      //untuk membuat server dan handle request HTTP
	"os"            //untuk ambil environtment variable
	"strconv"       //konversi tipe data
	"strings"       //manipulasi string
)

// struktur data produk
type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

// data sementara
var produk = []Produk{
	{
		ID:    1,
		Nama:  "Indomie Godog",
		Harga: 3500,
		Stok:  10,
	},
	{
		ID:    2,
		Nama:  "Vit 1000ml",
		Harga: 3000,
		Stok:  40,
	},
	{
		ID:    3,
		Nama:  "kecap",
		Harga: 12000,
		Stok:  20},
}

// fungsi helper untuk mengambil 1 produk berdasarkan ID
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	//ambil ID dari URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//jika user ngetik ID bukan angka
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	//cari produk didalam slice (array)
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	//jika loop selesai tapi ngga ketemu produknya
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

/*
fungsi untuk update data produk yang sudah ada (PUT)
PUT localhost:8080/api/produk/{id}
*/
func updateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	//baca data JSON baru yang dikirim user di Body request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//cari produk berdasarkan ID, lalu timpa datanya
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// fungsi DELETE
func deleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	//cari index produk yang mau dihapus
	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func main() {
	//route 1: endpoint ID (GET, PUT, DELETE)
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProdukByID(w, r)
		case "PUT":
			updateProduk(w, r)
		case "DELETE":
			deleteProduk(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//route 2: endpoint tanpa ID (GET ALL, POST)
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			//tampilkan semua data
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		case "POST":
			//tambah data baru
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			//auto increment ID
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) //kode 201 created
			json.NewEncoder(w).Encode(produkBaru)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//route health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	//setup port dynamic
	port := os.Getenv("PORT")
	if port == "" {
		port = "7860"
	}

	fmt.Println("Server running di port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
