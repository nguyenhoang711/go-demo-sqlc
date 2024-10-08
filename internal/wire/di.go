//go:build wireinject
package wire

import (
	"github.com/google/wire"
	"github.com/hoangnguyen/demo-sqlc/internal/controller"
	"github.com/hoangnguyen/demo-sqlc/internal/repo"
	"github.com/hoangnguyen/demo-sqlc/internal/service"
)

func InitUserController() (*controller.UserController, error) {
	wire.Build(
		repo.NewUSerRepo,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}