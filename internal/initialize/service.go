package initialize

import (
	"github.com/hoangnguyen/demo-sqlc/global"
	model "github.com/hoangnguyen/demo-sqlc/internal/models"
	"github.com/hoangnguyen/demo-sqlc/internal/service"
	"github.com/hoangnguyen/demo-sqlc/internal/service/impl"
)

func InitServiceInterface() {
	queries := model.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	service.InitUserAdmin(impl.NewUserAdminImpl(queries))

	global.Logger.Info("Init all service interface success!")
}