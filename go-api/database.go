package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"

    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Load environment variables dari .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("❌ Error loading .env file")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Gagal koneksi ke database:", err)
    }

    DB = database
    log.Println("✅ Koneksi ke MySQL berhasil")
}
