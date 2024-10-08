package service

import (
	"github.com/hoangnguyen/demo-sqlc/internal/repo"
	"github.com/hoangnguyen/demo-sqlc/pkg/response"
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
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. check email exists
	if us.userRepo.GetUserByGmail(email) {
		return response.ErrCodeUserEmailExist
	}
	return response.ErrCodeSuccess
}