package impl

import (
	"context"

	model "github.com/hoangnguyen/demo-sqlc/internal/models"
)

type sUserAdmin struct {
	re *model.Queries
}

func NewUserAdminImpl(r *model.Queries) *sUserAdmin {
	return &sUserAdmin{
		re: r,
	}
}

// implement the IUserLogin interface here
func (s *sUserAdmin) RemoveUser(ctx context.Context) error {
	return nil
}

func (s *sUserAdmin) FindOneUser(ctx context.Context) error {
	return nil
}