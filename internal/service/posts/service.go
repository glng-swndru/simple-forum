package posts

import (
	"context"

	"github.com/glng-swndru/simple-forum/internal/configs"
	"github.com/glng-swndru/simple-forum/internal/model/posts"
)

type PostRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error

	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error)
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
