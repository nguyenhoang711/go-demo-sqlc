package global

import (
	"github.com/hoangnguyen/demo-sqlc/pkg/logger"
	"github.com/hoangnguyen/demo-sqlc/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)