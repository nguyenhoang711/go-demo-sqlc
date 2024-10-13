package repo

import (
	"fmt"
	"time"

	"github.com/hoangnguyen/demo-sqlc/global"
)

type IUSerAuthRepository interface {
	AddOTP(email string, otp int, expiration int64) error
}

type userAuthRepository struct{}

// AddOTP implements IUSerAuthRepository.
func (u *userAuthRepository) AddOTP(email string, otp int, expiration int64) error {
	key := fmt.Sprintf("usr:%s:otp", email) // usr:email:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expiration)).Err()
}

func NewUserAuthRepository() IUSerAuthRepository {
	return &userAuthRepository{}
}
