package service

import "github.com/hoangnguyen/demo-sqlc/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		repo.NewUserRepo(),
	}
}


func (us *UserService) GetUserById(id string) string {
	return us.userRepo.GetUserDetail(id)
}