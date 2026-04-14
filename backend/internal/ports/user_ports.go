package ports

import (
	"antfarm/backend/internal/domain"
	"context"
)

type UserRepository interface {
	SaveUser(ctx context.Context, u *domain.User) error
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}
