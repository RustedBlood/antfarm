package application

import (
	"antfarm/backend/internal/domain"
	"antfarm/backend/internal/ports"
	"context"
)

// UserUseCase handles user-related business logic.
type UserUseCase struct {
	userRepo ports.UserRepository
}

func (uc *UserUseCase) SaveUserInfo(ctx context.Context, u *domain.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := uc.userRepo.SaveUser(ctx, u); err != nil {
		return err
	}

	return nil
}
