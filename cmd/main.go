package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/configs"
	"github.com/glng-swndru/simple-forum/internal/handlers/memberships"
	membershipRepo "github.com/glng-swndru/simple-forum/internal/repository/memberships"
	"github.com/glng-swndru/simple-forum/pkg/internalsql"
)

func main() {
	// Membuat instance default dari router Gin.
	r := gin.Default()

	var cfg *configs.Config // Variabel untuk menyimpan konfigurasi.

	// Inisialisasi konfigurasi aplikasi.
	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}), // Tentukan folder konfigurasi.
		configs.WithConfigFile("config"),                         // Nama file konfigurasi.
		configs.WithConfigType("yaml"),                           // Tipe file konfigurasi.
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config") // Hentikan aplikasi jika konfigurasi gagal di-load.
	}

	cfg = configs.Get() // Ambil konfigurasi yang sudah diinisialisasi.
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.Host)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	_ = membershipRepo.NewRepository(db)
	// Buat handler untuk route membership dan daftarkan route-nya.
	membershipHandler := memberships.NewHandler(r, nil)
	membershipHandler.RegisterRoute()

	// Jalankan server pada port yang ditentukan di file konfigurasi.
	r.Run(cfg.Service.Port)
}
