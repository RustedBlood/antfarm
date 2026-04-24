package userservice

import (
	db "antfarm/backend/internal/adapters/db/gen"
	"antfarm/backend/internal/domain/user"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}
func (ur *UserRepository) Save(ctx context.Context, u *user.User) error {
	track := pgtype.Text{
		String: u.TrackId,
		Valid:  true,
	}
	err := ur.q.SaveUser(ctx, db.SaveUserParams{
		ID:      u.Id,
		Name:    u.Name,
		Email:   u.Email,
		Phone:   u.Phone,
		TrackID: track,
	})
	return err
}
