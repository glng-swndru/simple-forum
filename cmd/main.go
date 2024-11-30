package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/configs"
	"github.com/glng-swndru/simple-forum/internal/handlers/memberships"
	"github.com/glng-swndru/simple-forum/internal/handlers/posts"
	membershipRepo "github.com/glng-swndru/simple-forum/internal/repository/memberships"
	postRepo "github.com/glng-swndru/simple-forum/internal/repository/posts"
	membershipSvc "github.com/glng-swndru/simple-forum/internal/service/memberships"
	postSvc "github.com/glng-swndru/simple-forum/internal/service/posts"

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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
