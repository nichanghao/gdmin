package cache

import (
	"context"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var (
	SysUserCache = new(sysUserCache)
)

const (
	SysUserStatusKey = "sys:user:status:"
)

type sysUserCache struct{}

func (*sysUserCache) SetSysUserStatus(userId uint64, status uint8) {

	ctx := context.Background()
	global.RedisCli.Set(ctx, SysUserStatusKey+strconv.FormatUint(userId, 10), status, time.Hour*24*7)
}

func (cache *sysUserCache) GetSysUserStatus(userId uint64) (status uint8, err error) {
	ctx := context.Background()
	result, _ := global.RedisCli.Get(ctx, SysUserStatusKey+strconv.FormatUint(userId, 10)).Result()
	if result != "" {
		if i, _ := strconv.ParseInt(result, 10, 8); i > 0 {
			return uint8(i), nil
		}
	}

	// 从数据库中获取状态数据
	var user model.SysUser
	if err = global.GormDB.Model(&model.SysUser{}).Where("id = ?", userId).Select("status").First(&user).Error; err != nil {
		zap.L().Error("Get sys user status from db error: ", zap.Error(err))
		return 0, buserr.NewNoticeBusErr("网络错误，请稍后再试！")
	} else {
		// 重新缓存数据
		cache.SetSysUserStatus(userId, user.Status)
		return user.Status, nil
	}
}
