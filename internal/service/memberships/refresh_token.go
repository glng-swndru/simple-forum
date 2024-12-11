package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/glng-swndru/simple-forum/internal/model/memberships"
	"github.com/glng-swndru/simple-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingResfreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err
	}

	if existingResfreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	// means the token in database is not matched with request token, throw error invalid refersh token
	if existingResfreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token has invalid")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exists")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}
	return token, nil
}
