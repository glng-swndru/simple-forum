package posts

import (
	"context"

	"github.com/glng-swndru/simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, PostID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, PostID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, PostID)
	if err != nil {
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, PostID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment to database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}