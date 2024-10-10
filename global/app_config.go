package global

import (
	"github.com/hoangnguyen/demo-sqlc/pkg/logger"
	"github.com/hoangnguyen/demo-sqlc/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
	Rdb *redis.Client
)