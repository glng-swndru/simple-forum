package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/configs"
	"github.com/glng-swndru/simple-forum/internal/handlers/memberships"
	membershipRepo "github.com/glng-swndru/simple-forum/internal/repository/memberships"
	membershipSvc "github.com/glng-swndru/simple-forum/internal/service/memberships"
	"github.com/glng-swndru/simple-forum/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config")
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.Host)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(membershipRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
