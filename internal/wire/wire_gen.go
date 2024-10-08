// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/hoangnguyen/demo-sqlc/internal/controller"
	"github.com/hoangnguyen/demo-sqlc/internal/repo"
	"github.com/hoangnguyen/demo-sqlc/internal/service"
)

// Injectors from di.go:

func InitUserController() (*controller.UserController, error) {
	iUserRepository := repo.NewUSerRepo()
	iUserService := service.NewUserService(iUserRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
