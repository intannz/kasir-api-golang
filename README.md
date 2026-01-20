# Simple Cashier API (Tugas 1 - Golang)

Ini adalah RESTful API sederhana untuk manajemen produk kasir. Dibuat menggunakan **Golang** (Go) murni tanpa framework pihak ketiga, hanya menggunakan library standar `net/http`.

Project ini dibuat untuk memenuhi tugas pemrograman backend menggunakan Golang.

## 🚀 Fitur

* **Create:** Menambahkan data produk baru.
* **Read:** Melihat semua produk atau satu produk spesifik.
* **Update:** Mengubah data produk (Nama, Harga, Stok).
* **Delete:** Menghapus data produk.

## 🛠️ Teknologi

* Golang (v1.20+)
* Standard Library (`net/http`, `encoding/json`)

## 📦 Cara Menjalankan (Run)

1.  Clone repository ini:
    ```bash
    git clone [https://github.com/intannz/kasir-api-golang.git](https://github.com/intannz/kasir-api-golang.git)
    ```
2.  Masuk ke folder project:
    ```bash
    cd nama-folder
    ```
3.  Jalankan server:
    ```bash
    go run main.go
    ```
4.  Server akan berjalan di `http://localhost:8080`.

## 🔗 Dokumentasi API

Gunakan Postman atau cURL untuk mengetes endpoint berikut:

| Method | Endpoint | Deskripsi | Body Request (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/produk` | Ambil semua data produk | - |
| `GET` | `/api/produk/{id}` | Ambil 1 produk (cth: `/api/produk/1`) | - |
| `POST` | `/api/produk` | Tambah produk baru | `{ "nama": "Barang X", "harga": 5000, "stok": 10 }` |
| `PUT` | `/api/produk/{id}` | Update data produk | `{ "nama": "Barang X Edit", "harga": 6000, "stok": 10 }` |
| `DELETE`| `/api/produk/{id}` | Hapus produk | - |
| `GET` | `/health` | Cek status server | - |

## 📝 Catatan

* Data disimpan sementara di memori (Slice), data akan reset jika server dimatikan.
* Port server bersifat dinamis (mengikuti environment variable `PORT`) atau default ke `8080` jika dijalankan di lokal.

---
**Happy Building! 🚀**