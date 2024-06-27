package initialize

import (
	"gitee.com/nichanghao/gdmin/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"os"
	"time"
)

// InitCasbin 初始化权限控制组件casbin
func InitCasbin() {
	// 加载 rbac 模型文件 https://github.com/casbin/casbin/blob/master/examples/rbac_model.conf
	rbacModel, err := model.NewModelFromFile("./rbac_model.conf")
	if err != nil {
		zap.L().Error("加载 rbac 模型文件失败：", zap.Error(err))
		os.Exit(1)
	}

	dbAdapter, err := gormAdapter.NewAdapterByDB(global.GormDB)
	if err != nil {
		zap.L().Error("连接数据库失败：", zap.Error(err))
		os.Exit(1)
	}

	// 创建一个可缓存的执行器
	enforcer, err := casbin.NewCachedEnforcer(rbacModel, dbAdapter)
	if err != nil {
		zap.L().Error("创建 casbin enforcer 失败：", zap.Error(err))
		os.Exit(1)
	}
	enforcer.SetExpireTime(time.Hour)

	if err := enforcer.LoadPolicy(); err != nil {
		zap.L().Error("加载权限策略失败：", zap.Error(err))
		os.Exit(1)
	}
	global.Enforcer = enforcer

}
