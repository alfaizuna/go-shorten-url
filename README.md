# URL Shortener Sederhana (Go)

Layanan ini berfungsi untuk memperpendek URL seperti bit.ly, dibuat dengan bahasa Go menggunakan arsitektur yang bersih dan modular.

## Struktur Project

```
urlshortener/
├── main.go              # Entry point aplikasi
├── config/
│   └── config.go        # Konfigurasi aplikasi
├── handlers/
│   └── url_handler.go   # HTTP handlers
├── storage/
│   └── storage.go       # Interface dan implementasi storage
├── utils/
│   └── encoding.go      # Fungsi encoding URL
└── README.md
```

## Fitur
- **POST /shorten**: Memasukkan URL panjang, mengembalikan URL pendek.
- **GET /{shortPath}**: Redirect ke URL asli berdasarkan path pendek.
- Penyimpanan in-memory (map) atau Redis (jika tersedia).
- Encoding otomatis shortPath dari URL panjang.
- Konfigurasi fleksibel melalui environment variables.

## Cara Menjalankan

### 1. Install Dependensi
```bash
go mod tidy
```

### 2. Jalankan Server (Default)
```bash
go run main.go
```

### 3. (Opsional) Gunakan Redis
```bash
export REDIS_ADDR=localhost:6379
export BASE_URL=http://localhost:8080
export PORT=8080
go run main.go
```

## Environment Variables

| Variable | Default | Deskripsi |
|----------|---------|-----------|
| `PORT` | `8080` | Port server |
| `REDIS_ADDR` | - | Address Redis (opsional) |
| `BASE_URL` | `http://localhost:8080` | Base URL untuk short URL |

## Contoh Penggunaan

### 1. Memperpendek URL
```bash
curl -X POST -d "url=https://www.google.com" http://localhost:8080/shorten
```

Response:
```
http://localhost:8080/abc123
```

### 2. Redirect
Buka hasil short URL di browser:
```
http://localhost:8080/abc123
```
Akan otomatis redirect ke `https://www.google.com`.

## Arsitektur

### Separation of Concerns
- **Handlers**: Menangani HTTP requests dan responses
- **Storage**: Interface dan implementasi penyimpanan (Map/Redis)
- **Utils**: Fungsi helper seperti encoding
- **Config**: Manajemen konfigurasi aplikasi

### Dependency Injection
- Storage diinjeksi ke handler
- Config diinjeksi ke handler dan storage
- Memudahkan testing dan maintenance

## Catatan
- ShortPath dihasilkan otomatis dari hash URL panjang (6 karakter base64).
- Tidak ada validasi duplikasi, jika URL sama akan menghasilkan shortPath yang sama.
- Untuk produksi, sebaiknya tambahkan validasi, logging, dan fitur keamanan.

---

Dibuat untuk latihan mini project Golang dengan arsitektur profesional. 