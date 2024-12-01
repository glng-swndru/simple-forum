package posts

import (
	"context"

	"github.com/glng-swndru/simple-forum/internal/configs"
	"github.com/glng-swndru/simple-forum/internal/model/posts"
)

type PostRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
}

type service struct {
	cfg      *configs.Config
	postRepo PostRepository
}

func NewService(cfg *configs.Config, postRepo PostRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
