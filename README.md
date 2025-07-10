# URL Shortener Sederhana (Go)

Layanan ini berfungsi untuk memperpendek URL seperti bit.ly, dibuat dengan bahasa Go.

## Fitur
- **POST /shorten**: Memasukkan URL panjang, mengembalikan URL pendek.
- **GET /{shortPath}**: Redirect ke URL asli berdasarkan path pendek.
- Penyimpanan in-memory (map) atau Redis (jika tersedia).
- Encoding otomatis shortPath dari URL panjang.

## Cara Menjalankan
1. **Clone repo & install dependensi**
   ```bash
   go mod tidy
   ```
2. **Jalankan server**
   ```bash
   go run main.go
   ```
3. **(Opsional) Gunakan Redis**
   - Jalankan Redis server.
   - Set environment variable sebelum menjalankan:
     ```bash
     export REDIS_ADDR=localhost:6379
     go run main.go
     ```

## Contoh Penggunaan
### 1. Memperpendek URL
Request:
```bash
curl -X POST -d "url=https://www.google.com" http://localhost:8080/shorten
```
Response:
```
http://localhost:8080/abc123 (contoh)
```

### 2. Redirect
Buka hasil short URL di browser:
```
http://localhost:8080/abc123
```
Akan otomatis redirect ke `https://www.google.com`.

## Struktur Kode
- `main.go` : Seluruh logic utama (routing, encoding, penyimpanan, handler)
- Penyimpanan otomatis menggunakan Redis jika `REDIS_ADDR` di-set, jika tidak maka pakai map (in-memory).

## Catatan
- ShortPath dihasilkan otomatis dari hash URL panjang (6 karakter base64).
- Tidak ada validasi duplikasi, jika URL sama akan menghasilkan shortPath yang sama.
- Untuk produksi, sebaiknya tambahkan validasi, logging, dan fitur keamanan.

---

Dibuat untuk latihan mini project Golang. 