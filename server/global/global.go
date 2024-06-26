package global

import (
	"gitee.com/nichanghao/gdmin/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Config 全局配置
	Config *config.Config

	// LOG 全局日志
	LOG *zap.Logger

	// GormDB 全局数据库连接
	GormDB *gorm.DB

	RedisCli *redis.Client
)
