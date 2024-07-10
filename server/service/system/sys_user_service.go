package system

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/utils"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
)

type SysUserService struct {
}

// Login 用户登录
func (userService *SysUserService) Login(user *model.SysUser) (*response.SysUserLoginResp, error) {

	var userRes *model.SysUser

	if err := global.GormDB.Where("username = ?", user.Username).First(&userRes).Error; err != nil {
		return nil, common.NewNoticeBusErr("用户名或密码错误！")
	}

	// 校验密码
	if ok := utils.BCRYPT.CheckPasswordHash(user.Password, userRes.Password); !ok {
		return nil, common.NewNoticeBusErr("用户名或密码错误！")
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

// PageUsers 分页查询用户列表
func (userService *SysUserService) PageUsers(req *request.SysUserPageReq) ([]*model.SysUser, error) {
	tx := global.GormDB.Model(&model.SysUser{}).Limit(req.Limit).Offset(req.Offset)
	if req.Username != "" {
		tx.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Nickname != "" {
		tx.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}
	if req.Phone != "" {
		tx.Where("phone LIKE ?", "%"+req.Phone+"%")
	}
	if req.Email != "" {
		tx.Where("email LIKE ?", "%"+req.Email+"%")
	}
	var userList []*model.SysUser

	if err := tx.Find(&userList).Error; err != nil {
		return userList, err
	}

	return userList, nil
}

// EditUser 编辑用户
func (userService *SysUserService) EditUser(user *model.SysUser) error {

	if err := global.GormDB.Model(user).Select("nickname", "phone", "email").Updates(user).Error; err != nil {
		return err
	}

	return nil
}
