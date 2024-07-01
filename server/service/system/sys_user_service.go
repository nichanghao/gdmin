package system

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/utils"
	"gitee.com/nichanghao/gdmin/web/response"
)

type SysUserService struct {
}

// Login 用户登录
func (userService *SysUserService) Login(user *model.SysUser) (*response.SysUserLoginResp, error) {

	var userRes *model.SysUser

	if err := global.GormDB.Where("username = ?", user.Username).First(&userRes).Error; err != nil {
		return nil, err
	}

	// 校验密码
	if ok := utils.BCRYPT.CheckPasswordHash(user.Password, userRes.Password); !ok {
		return nil, common.ErrPassWdNonMatched
	}

	// 生成token
	token, err := utils.JWT.GenerateToken(&common.UserClaims{
		ID: userRes.Id, Username: userRes.Username, NickName: userRes.Nickname,
	})
	if err != nil {
		return nil, err
	}

	return &response.SysUserLoginResp{Token: token, UserInfo: userRes}, nil
}
