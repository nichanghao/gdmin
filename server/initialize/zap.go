package initialize

import (
	"gitee.com/nichanghao/gdmin/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitZap 初始化日志组件
func InitZap() {

	zapCfg := global.Config.Zap

	// 配置 Lumberjack 日志轮转
	lumberjackLogger := &lumberjack.Logger{
		Filename:   zapCfg.Filename,
		MaxSize:    zapCfg.MaxSize,    // 单个文件最大大小 100M
		MaxBackups: zapCfg.MaxBackups, // 保留旧日志文件个数
		MaxAge:     zapCfg.MaxAge,     // 日志保留的天数
		Compress:   zapCfg.Compress,   // 是否压缩/归档旧日志文件
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 创建 Core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberjackLogger),
		// 获取日志级别
		parseZapLevel(zapCfg.Level),
	)

	// 创建 Logger
	logger := zap.New(core)

	// 替换zap日志对象
	zap.ReplaceGlobals(logger)
}

func parseZapLevel(logLevel string) (level zapcore.Level) {
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	return
}
