## About

dot-test adalah API server untuk tech test freelancer DOT.
Aplikasi ini menggunakan bahasa pemrograman GO, framework Echo dan GORM.

### Struktur Folder

Struktur folder dot-test ini mengadopsi struktur folder pada Ruby on Rails. Untuk menjalankan aplikasi, gunakan perintah
```
$ go run bin/main.go
```

### Clean Architecture

Selain itu dot-test ini menggunakan clean architecture yang dijelaskan oleh Uncle Bob (Referensi: https://medium.com/golangid/mencoba-golang-clean-architecture-c2462f355f41)
dengan 4 layer pada arsitektur, yaitu :
1. Models
2. Repository
3. Usecase
4. Delivery

Arsitektur tersebut saya tempatkan pada path `app/pkg/`.

### Error Handler

Sedangkan untuk error handler, saya tempatkan log history-nya pada path `logs`.

### Redis Cache

Penamaan cache key saya kelompokkan berdasarkan jenis data. Tiap jenis data akan menyimpan cache berisi daftar semua data dan spesifik data dengan contoh format
- Semua data : `dot-book`
- Spesifik data : `dot-book_{{UUID}}`

Hal ini bertujuan untuk memudahkan pemanggilan data.

### End-to-end Test

End-to-end test ditempatkan pada path `app/pkg/{{jenis_data}}`.

## Requirement

- Go 1.14
- PostgreSQL
- Redis
- Docker

## How to run

1. Create .env and copy variables from .env.example

2. Running app
```
$ go run bin/main.go
```