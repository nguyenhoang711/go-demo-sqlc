package logger

import (
	"os"

	"github.com/hoangnguyen/demo-sqlc/pkg/settings"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}


func NewLogger(config settings.LoggerSetting) *LoggerZap{
	logLevel := config.Log_level //mac dinh
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "panic":
		level = zapcore.PanicLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name, //"./storages/logs/dev.xxx.log",
		MaxSize:    config.Max_size,      // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)
	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	//1716714967.877995 -> 2024-05-26T16:18:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts -> time
	// encodeConfig.TimeKey = "time"

	// from info --> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller": "cli/main.log.go:24"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}