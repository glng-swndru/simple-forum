package memberships

import (
	"context"

	"github.com/glng-swndru/simple-forum/internal/model/memberships"
)

type MembershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	membershipRepo MembershipRepository
}

func NewService(membershipRepo MembershipRepository) *service {
	return &service{
		membershipRepo: membershipRepo,
	}
}
