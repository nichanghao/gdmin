package global

import (
	"gitee.com/nichanghao/gdmin/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// Config 全局配置
	Config *config.Config

	// GormDB 全局数据库连接
	GormDB *gorm.DB

	RedisCli *redis.Client
)
