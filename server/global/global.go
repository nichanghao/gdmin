package global

import (
	"gitee.com/nichanghao/gdmin/config"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// Config 全局配置
	Config *config.Config

	GinEngine *gin.Engine

	// GormDB 全局数据库连接
	GormDB *gorm.DB

	// RedisCli Redis客户端
	RedisCli *redis.Client

	// Enforcer 全局Casbin权限管理器
	Enforcer *casbin.CachedEnforcer

	// PermissionRouter 权限路由表
	PermissionRouter = make(map[string]string, 256)
)
