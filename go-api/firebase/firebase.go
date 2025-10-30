package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

// InitFirebase digunakan untuk menginisialisasi koneksi Firebase
func InitFirebase() {
	opt := option.WithCredentialsFile("go-api/firebase/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("❌ Gagal menginisialisasi Firebase: %v", err)
	}
	App = app
	log.Println("✅ Firebase berhasil diinisialisasi")
}
