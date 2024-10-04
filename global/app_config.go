package global

import (
	"github.com/hoangnguyen/demo-sqlc/pkg/logger"
	"github.com/hoangnguyen/demo-sqlc/pkg/settings"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
)