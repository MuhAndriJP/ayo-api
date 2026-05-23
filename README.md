# AYO API — Backend Technical Test

REST API untuk manajemen tim sepakbola amatir perusahaan XYZ.  
Stack: **Go + Gin + GORM + MySQL** | Auth: **JWT (HS256) + bcrypt**

---

## Prasyarat

- Go 1.21+
- MySQL 8.0+

---

## Setup

### 1. Clone & masuk ke folder
```bash
cd /path/to/ayo
```

### 2. Buat database MySQL
```sql
CREATE DATABASE ayo_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. Konfigurasi environment
```bash
cp .env.example .env
# Edit .env — isi DB_USER, DB_PASSWORD, DB_NAME, JWT_SECRET
```

### 4. Install dependencies
```bash
go mod tidy
```

### 5. Jalankan server
```bash
go run cmd/server/main.go
# Server running on :8080
```

---

## Menjalankan Tests

```bash
go test ./tests/... -v
```

---

## Struktur Project

```
docs/ayo-api.postman_collection   — REST client collection
docs/schema.dbmlk                 — Database schema
cmd/server/main.go                — entry point, dependency wiring
internal/
  config/                         — load .env
  database/                       — GORM connect + AutoMigrate
  model/                          — GORM models (soft delete via DeletedAt)
  dto/                            — request/response structs
  repository/                     — data access layer (interface-based)
  service/                        — business logic
  handler/                        — HTTP handlers (Gin)
  middleware/                     — JWT auth, rate limiter
  router/                         — route registration
  util/                           — bcrypt, JWT, response envelope
tests/service/                    — unit tests (mock repository)
uploads/teams/                    — penyimpanan logo (gitignored)
```

---

## Daftar Endpoint

Base URL: `http://localhost:8080/api/v1`

Semua endpoint kecuali `/auth/*` memerlukan header:
```
Authorization: Bearer <token>
```

### Auth

| Method | Path | Keterangan |
|--------|------|-----------|
| POST | `/auth/register` | Daftar admin |
| POST | `/auth/login` | Login → JWT 24 jam |

### Teams

| Method | Path | Keterangan |
|--------|------|-----------|
| GET | `/teams` | List tim |
| GET | `/teams/:id` | Detail tim |
| POST | `/teams` | Buat tim baru |
| PUT | `/teams/:id` | Update tim |
| DELETE | `/teams/:id` | Soft delete tim |

### Players

| Method | Path | Keterangan |
|--------|------|-----------|
| GET | `/teams/:teamId/players` | List pemain dalam tim |
| POST | `/teams/:teamId/players` | Tambah pemain ke tim |
| GET | `/players/:id` | Detail pemain |
| PUT | `/players/:id` | Update pemain |
| DELETE | `/players/:id` | Soft delete pemain |

### Matches

| Method | Path | Keterangan |
|--------|------|-----------|
| GET | `/matches` | List pertandingan |
| GET | `/matches/:id` | Detail pertandingan + gol |
| POST | `/matches` | Buat jadwal pertandingan |
| PUT | `/matches/:id` | Update jadwal |
| DELETE | `/matches/:id` | Soft delete pertandingan |
| POST | `/matches/:id/result` | Laporkan hasil + gol |
| PUT | `/matches/:id/result` | Koreksi hasil + gol |
| GET | `/matches/:id/report` | Report lengkap pertandingan |

### Lainnya

| Method | Path | Keterangan |
|--------|------|-----------|
| GET | `/health` | Health check |
| GET | `/uploads/teams/:filename` | Serve logo tim |

---

## Format Respons

Semua respons menggunakan envelope JSON:

```json
{
  "success": true,
  "message": "berhasil",
  "data": { ... }
}
```

Error:
```json
{
  "success": false,
  "message": "pesan error",
  "errors": "detail error"
}
```

---

## Fitur Keamanan

- Password di-hash dengan **bcrypt** (cost 12). Hash tidak pernah dikembalikan ke client.
- JWT HS256, expire 24 jam. Secret dari env variable.
- Upload logo: validasi MIME (JPEG/PNG only), max 2MB, disimpan dengan nama UUID.
- Rate limit endpoint login: 10 request per 15 menit per IP.
- Soft delete: record dihapus secara logis (`deleted_at IS NOT NULL`), data tetap ada di database.
- Nomor punggung unik per tim: dicek di service layer sebelum insert/update.
- Input validation: semua request body/query divalidasi via struct tags.

---

## Asumsi & Keputusan Desain

1. **Single admin role.** Tidak ada role player/coach — semua operasi CRUD dilakukan oleh admin.
2. **Logo disimpan lokal** di folder `uploads/` (tidak ke S3). Dapat diswap via interface storage tanpa mengubah service layer.
3. **Soft delete tim tidak otomatis menghapus pemain.** Pemain yang timnya dihapus tetap ada di database dan bisa di-delete secara manual. Ini disengaja agar tidak ada data gol/pertandingan yang rusak.
4. **AutoMigrate GORM** digunakan untuk kemudahan setup. Di production, disarankan menggunakan migration tool seperti golang-migrate.
5. **Akumulasi kemenangan** dihitung dari semua pertandingan `finished` dengan `match_date <= tanggal pertandingan terkait`. Draw tidak dihitung sebagai kemenangan.
6. **Tanggal dan waktu pertandingan** disimpan terpisah sesuai spesifikasi soal.