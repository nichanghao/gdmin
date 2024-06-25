package system

import (
	"errors"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/util"
)

type SysUserService struct {
}

// Login 用户登录
func (userService *SysUserService) Login(user *model.SysUser) (UserRes *model.SysUser, err error) {

	if err = global.GormDB.Where("username = ?", user.Username).First(UserRes).Error; err != nil {
		return nil, err
	}

	if ok := util.CheckPasswordHash(user.Password, UserRes.Password); !ok {
		return nil, errors.New("密码错误！")
	}

	return UserRes, nil
}
