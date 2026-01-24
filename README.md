# Simple Cashier API (Tugas 1 - Golang)

Repo ini adalah RESTful API sederhana untuk manajemen produk dan kategori kasir. Dibuat menggunakan **Golang** (Go) murni tanpa framework pihak ketiga, hanya menggunakan library standar `net/http` dan **Swagger** untuk dokumentasi.

Project ini dibuat untuk memenuhi tugas pemrograman backend menggunakan Golang.

## 🌐 Live Demo & Dokumentasi
Aplikasi ini sudah di-deploy menggunakan **Zeabur** dan memiliki dokumentasi lengkap via **Swagger UI**.

👉 **Akses Dokumentasi API (Swagger):**
`https://kasir-api-toko.zeabur.app/swagger/index.html`


## 🚀 Fitur

* **Manajemen Produk:** Create, Read, Update, Delete (CRUD) data produk.
* **Manajemen Kategori:** Create, Read, Update, Delete (CRUD) data kategori.
* **In-Memory Storage:** Penyimpanan data sementara menggunakan Slice/Array.
* **API Documentation:** Dokumentasi otomatis menggunakan Swagger.

## 🛠️ Teknologi

* **Golang** (v1.20+)
* **Standard Library** (`net/http`, `encoding/json`)
* **Swaggo** (Untuk generate Swagger Docs)
* **Zeabur** (Deployment Platform)

## 📦 Cara Menjalankan (Local)

1.  Clone repository ini:
    ```bash
    git clone https://github.com/intannz/kasir-api-golang.git
    ```
2.  Masuk ke folder project:
    ```bash
    cd kasir-api-golang
    ```
3.  Jalankan server:
    ```bash
    go run main.go
    ```
4.  Buka Swagger di browser:
    `http://localhost:8080/swagger/index.html`

## 🔗 Daftar Endpoint Utama

Gunakan Swagger UI untuk pengetesan yang lebih mudah, atau gunakan Postman/cURL:

### 🛒 Products
| Method | Endpoint | Deskripsi | Contoh Body Request (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/products` | Ambil semua produk | - |
| `GET` | `/api/products/{id}` | Ambil 1 produk | - |
| `POST` | `/api/products` | Tambah produk | `{ "name": "Latte", "price": 18000, "stock": 20 }` |
| `PUT` | `/api/products/{id}` | Update produk | `{ "name": "Latte Edit", "price": 20000, "stock": 15 }` |
| `DELETE`| `/api/products/{id}` | Hapus produk | - |

### 🏷️ Categories
| Method | Endpoint | Deskripsi | Contoh Body Request (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/categories` | Ambil semua kategori | - |
| `POST` | `/categories` | Tambah kategori | `{ "name": "Minuman", "description": "Aneka Kopi" }` |
| `PUT` | `/categories/{id}` | Update kategori | `{ "name": "Beverages", "description": "Coffee & Tea" }` |
| `DELETE`| `/categories/{id}` | Hapus kategori | - |

## 📝 Catatan

* **Data Reset:** Karena menggunakan *In-Memory* (variabel slice), data akan kembali ke default jika server di-restart (deploy ulang).
* **Environment:** Server berjalan di port `8080` secara default.

---
**Happy Building! 🚀**