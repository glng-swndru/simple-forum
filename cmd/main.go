package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
