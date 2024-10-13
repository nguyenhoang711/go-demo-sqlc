package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hoangnguyen/demo-sqlc/global"
	"github.com/hoangnguyen/demo-sqlc/internal/repo"
	"github.com/hoangnguyen/demo-sqlc/internal/utils/random"
	"github.com/hoangnguyen/demo-sqlc/internal/utils/sendto"
	"github.com/hoangnguyen/demo-sqlc/pkg/response"
	"go.uber.org/zap"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetUserById(id string) string {
// 	return us.userRepo.GetUserDetail(id)
// }

// INTERFACE_VERSION
type IUserService interface {
	Register(email, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	userAuthRepo repo.IUSerAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUSerAuthRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hash email

	// 5. check OTP is available

	// 6. user spam ...

	// 1. check email exists
	if us.userRepo.GetUserByGmail(email) {
		return response.ErrCodeUserEmailExist
	}

	// 2. new OTP ...
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("Otp test is %d\n", otp)

	// 3. save OTP in Redis with expiration time
	err := us.userAuthRepo.AddOTP(email, otp, int64(10 * time.Minute))
	if err != nil {
		return response.ErrIOTPInvalid
	}
	
	// 4. send Email OTP
	// err = sendto.SendTextEmailOtp([]string{email}, "phuocnt.hutech@gmail.com", strconv.Itoa(otp))
	err = sendto.SendTemplateEmailOtp([]string{email}, "phuocnt.hutech@gmail.com", "otp-auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	if err != nil {
		global.Logger.Error("Error sending email otp::", zap.Error(err))
		return response.ErrSendEmailOTPFail
	}
	
	return response.ErrCodeSuccess
}