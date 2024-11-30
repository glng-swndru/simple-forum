package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/middleware"
	"github.com/glng-swndru/simple-forum/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, UserID int64, req posts.CreatePostRequest) error
}
type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
}
