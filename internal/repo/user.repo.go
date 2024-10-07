package repo

import "fmt"

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func(ur *UserRepo) GetUserDetail(id string) string {
	return fmt.Sprintf("Tips %s", id)
}