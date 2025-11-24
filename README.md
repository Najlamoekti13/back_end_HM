
# prosedural back-end

- documentation using swagger: http://localhost:8080/swagger/index.html
- endpoint_port:
        - 8080


## catatan penggunaan

GET → meminta data dari server

POST → mengirim data baru ke server

PUT / PATCH → mengubah data yang ada

DELETE → menghapus data

## Deskripsi
API ini dibuat menggunakan Go + Gin + GORM.
Tujuannya untuk belajar membuat backend untuk aplikasi Android.

## Struktur Folder
- main.go      : server & route
- database.go  : koneksi database
- models/      : tempat model
- docs/        : Swagger docs

## Endpoint
- GET /hello         → tes API
- GET /users         → ambil semua user
- POST /users        → tambah user baru
- PUT /users/:id     → update user
- DELETE /users/:id  → hapus user

NOTE: swagger digunakan untuk notification, yang lain nya menggunakan firestore emulator

## Catatan
- Pastikan MySQL jalan sebelum run server
- Gunakan `go run .` agar semua file main + database bisa dijalankan
