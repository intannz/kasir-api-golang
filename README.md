# Simple Cashier API

Repo ini adalah RESTful API sederhana untuk manajemen produk dan kategori kasir. Dibuat menggunakan **Golang** (Go) murni tanpa framework pihak ketiga, hanya menggunakan library standar `net/http` dan **Swagger** untuk dokumentasi.

Project ini dibuat untuk memenuhi tugas pemrograman backend menggunakan Golang.

## 🌐 Live Demo & Dokumentasi
Aplikasi ini sudah di-deploy menggunakan **Zeabur** dan memiliki dokumentasi lengkap via **Swagger UI**.

👉 **Akses Dokumentasi API (Swagger):**
[https://kasir-api-toko.zeabur.app/swagger/index.html](https://kasir-api-toko.zeabur.app/swagger/index.html)


## 🚀 Fitur

* **Database Storage:** Data disimpan secara permanen di **PostgreSQL** (via Supabase), tidak hilang saat server restart.
* **Framework Gin:** Routing dan handling request lebih cepat dan efisien menggunakan Gin Gonic.
* **Configuration Management:** Pengaturan environment (Database URL, Port) dikelola menggunakan **Viper**.
* **CRUD Lengkap:** Create, Read, Update, Delete untuk Produk dan Kategori.
* **API Documentation:** Dokumentasi interaktif otomatis dengan Swagger.

## 🛠️ Teknologi

* **Golang** (v1.20+)
* **Gin Gonic** (Web Framework)
* **PostgreSQL** (Database)
* **Supabase** (Cloud Database Provider)
* **Viper** (Config Management)
* **lib/pq** (Postgres Driver)
* **Swaggo** (Swagger Docs Generator)
* **Zeabur** (Deployment)

## ⚠️ Catatan Khusus (Supabase User)

Jika Anda menggunakan **Supabase Transaction Pooler** (Port 6543) dan mengalami error `binary_parameters` atau koneksi `EOF`, disarankan menggunakan driver `lib/pq` versi **v1.10.9**:
```bash
go get github.com/lib/pq@v1.10.9
```

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

---
**Happy Building! 🚀**