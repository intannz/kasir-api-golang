# Simple Cashier API

Repo ini adalah RESTful API sederhana untuk manajemen produk dan kategori kasir. Dibuat menggunakan **Golang** (Go) murni tanpa framework pihak ketiga, hanya menggunakan library standar `net/http` dan **Swagger** untuk dokumentasi.

Project ini dibuat untuk memenuhi tugas pemrograman backend menggunakan Golang.

## 🌐 Live Demo & Dokumentasi
Aplikasi ini sudah di-deploy menggunakan **Zeabur** dan memiliki dokumentasi lengkap via **Swagger UI**.

👉 **Akses Dokumentasi API (Swagger):**
[https://kasir-api-toko.zeabur.app/swagger/index.html](https://kasir-api-toko.zeabur.app/swagger/index.html)

## ⚠️ PENTING: Panduan Testing (Urutan Eksekusi)
Karena API ini menggunakan **Relational Database** (Product membutuhkan Category), harap ikuti urutan tes berikut agar tidak terjadi error:

### 1️⃣ Langkah 1: Buat Kategori (Category)
Anda **WAJIB** membuat minimal satu kategori terlebih dahulu sebagai "wadah" untuk produk.
* **Endpoint:** `POST /categories`
* **Contoh Body:**
    ```json
    {
      "name": "Makanan Berat",
      "description": "Nasi dan Lauk Pauk"
    }
    ```
* *Catat ID yang terbentuk (Misal: ID = 1).*

### 2️⃣ Langkah 2: Buat Produk (Product)
Setelah kategori ada, baru Anda bisa membuat produk. Pastikan `categoryId` diisi dengan ID kategori yang valid.
* **Endpoint:** `POST /api/products`
* **Contoh Body:**
    ```json
    {
      "name": "Nasi Goreng Spesial",
      "price": 25000,
      "stock": 50,
      "categoryId": 1
    }
    ```

### 3️⃣ Langkah 3: Cek Hasil Join (Explore Join)
Untuk melihat implementasi **SQL INNER JOIN**, panggil endpoint Get All Products. API akan mengembalikan data produk lengkap dengan **Nama Kategori**-nya.
* **Endpoint:** `GET /api/products`
* **Hasil Response:**
    ```json
    [
      {
        "id": 1,
        "name": "Nasi Goreng Spesial",
        "price": 25000,
        "categoryId": 1,
        "categoryName": "Makanan Berat"  <-- Data dari tabel Categories (Hasil Join)
      }
    ]
    ```
### 4️⃣ Langkah 4: Proses Checkout (Transaksi)
Setelah produk dibuat dan memiliki stok, Anda bisa mencoba fitur transaksi. Fitur ini sudah dilengkapi dengan **Database Transaction** (Anti data korup) dan **Row Locking** (Anti rebutan stok).
* **Endpoint:** `POST /api/checkout`
* **Contoh Body:**
    ```json
    {
      "items": [
        { "product_id": 1, "quantity": 2 }
      ]
    }
    ```
* *💡 Tips: Cek kembali endpoint `GET /api/products` setelah checkout berhasil. Stok produk otomatis berkurang!*

### 5️⃣ Langkah 5: Cek Laporan Penjualan (Report)
Setelah melakukan beberapa transaksi, Anda bisa melihat ringkasan omzet dan produk paling laku terjual.
* **Endpoint:** `GET /api/report` (Default: Hari Ini)
* **Contoh dengan Filter Tanggal:** `GET /api/report?start_date=2026-02-01&end_date=2026-02-28`

---

## 🚀 Fitur

* **Database Persistent:** Data disimpan aman di PostgreSQL (Supabase).
* **Relasi One-to-Many:** Satu Kategori memiliki banyak Produk.
* **SQL Join Query:** Endpoint produk menampilkan data gabungan dari tabel Kategori.
* **Environment Config:** Konfigurasi sensitif (DB URL) aman menggunakan `.env`.
* **Auto Docs:** Dokumentasi otomatis via Swagger.
* **Database Transaction (ACID):** Menjamin keamanan data saat *checkout*. Jika terjadi error di tengah proses, seluruh perubahan akan di-*rollback*.
* **Row-Level Locking (`FOR UPDATE`):** Mencegah *race condition* saat kasir memproses barang yang sama bersamaan, memastikan stok tidak pernah minus.
* **Search & Filter:** Pencarian produk berdasarkan nama menggunakan URL Query Parameter.
* **Data Aggregation:** Query SQL tingkat lanjut (`SUM`, `COUNT`, `GROUP BY`) untuk menghasilkan laporan penjualan.

## 🛠️ Teknologi

* **Golang** (Backend Logic)
* **PostgreSQL** (Database)
* **Lib/PQ** (Postgres Driver)
* **Swaggo** (API Documentation)
* **Zeabur** (Cloud Deployment)

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

Gunakan Swagger UI untuk pengetesan yang lebih mudah.

### 🏷️ Categories (Buat Ini dulu karena Products butuh categoryId)
| Method | Endpoint | Deskripsi | Contoh Body Request (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/categories` | Ambil semua kategori | - |
| `POST` | `/categories` | Tambah kategori | `{ "name": "Minuman", "description": "Aneka Kopi" }` |
| `PUT` | `/categories/{id}` | Update kategori | `{ "name": "Beverages", "description": "Coffee & Tea" }` |
| `DELETE`| `/categories/{id}` | Hapus kategori | - |

### 🛒 Products (Butuh categoryId)
| Method | Endpoint | Deskripsi | Contoh Body Request (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/products` | Ambil semua produk (Bisa tambah `?name=indomie` untuk pencarian) | - |
| `GET` | `/api/products/{id}` | Ambil 1 produk | - |
| `POST` | `/api/products` | Tambah produk | `{ "name": "Latte", "price": 18000, "stock": 20, "categoryId": 1 }` |
| `PUT` | `/api/products/{id}` | Update produk | `{ "name": "Latte Edit", "price": 20000, "stock": 15, "categoryId": 1 }` |
| `DELETE`| `/api/products/{id}` | Hapus produk | - |

### 🛍️ Transactions & Reports
| Method | Endpoint | Deskripsi | Contoh Body Request (JSON) |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/checkout` | Proses pembayaran & potong stok otomatis | `{ "items": [ { "product_id": 1, "quantity": 2 } ] }` |
| `GET` | `/api/report` | Laporan penjualan (omzet & produk terlaris) | - |
---
**Happy Building! 🚀**