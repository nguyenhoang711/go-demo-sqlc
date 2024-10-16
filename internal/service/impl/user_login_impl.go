package impl

import (
	"context"

	model "github.com/hoangnguyen/demo-sqlc/internal/models"
)

type sUserLogin struct {
	re *model.Queries
}

func NewUserLoginImpl(r *model.Queries) *sUserLogin {
	return &sUserLogin{
		re: r,
	}
}

// implement the IUserLogin interface here
func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
