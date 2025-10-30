package main

import (
    "log"

    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    _ "github.com/najla/go-api/docs"     // ubah sesuai nama module kamu
    "github.com/najla/go-api/firebase"  // 🔥 Firebase
)

// @title Example API
// @version 1.0
// @description Contoh REST API menggunakan Go + Firebase + Swagger
// @host localhost:8080
// @BasePath /

// HelloWorld godoc
// @Summary Menyapa pengguna
// @Description Mengembalikan pesan sapaan
// @Tags contoh
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func main() {
    // 1️⃣ Inisialisasi Firebase
    firebase.InitFirebase()

    // 2️⃣ Hubungkan ke database lokal (sementara dimatikan)
    // ConnectDatabase()

    // 3️⃣ Setup Gin
    r := gin.Default()

    // Swagger endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Endpoint sederhana
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "HealthMate_Back-end 🚀 Firebase connected successfully!",
        })
    })

    log.Println("🌐 Server berjalan di http://localhost:8080")
    r.Run(":8080")
}
