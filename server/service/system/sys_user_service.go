package system

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/utils"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"gorm.io/gorm"
)

type SysUserService struct {
}

// Login 用户登录
func (userService *SysUserService) Login(user *model.SysUser) (*response.SysUserLoginResp, error) {

	var userRes *model.SysUser

	if err := global.GormDB.Where("username = ?", user.Username).Preload("Roles").First(&userRes).Error; err != nil {
		return nil, buserr.NewNoticeBusErr("用户不存在！")
	}

	// 校验密码
	if ok := utils.BCRYPT.CheckPasswordHash(user.Password, userRes.Password); !ok {
		return nil, buserr.NewNoticeBusErr("用户名或密码错误！")
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

// ResetPassword 重置密码
func (userService *SysUserService) ResetPassword(req *request.SysUserEditReq) error {

	password, err := utils.BCRYPT.HashPassword(req.Password)
	if err != nil {
		return err
	}

	if err = global.GormDB.Model(&model.SysUser{}).Where("id = ?", req.Id).Update("password", password).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser 删除用户
func (userService *SysUserService) DeleteUser(id uint64) error {

	return global.GormDB.Transaction(func(tx *gorm.DB) error {

		// 删除用户并同步删除关联的角色
		if err := tx.Select("Roles").Delete(&model.SysUser{Id: id}).Error; err != nil {
			return err
		}

		// 删除casbin用户
		if err := CasbinService.ClearRolesForUser(id); err != nil {
			return err
		}

		return nil
	})
}

// AllocateRoles 分配角色给用户
func (userService *SysUserService) AllocateRoles(req *request.SysUserEditReq) error {

	var roles []*model.SysRole
	for i := range req.RoleIds {
		roles = append(roles, &model.SysRole{Id: req.RoleIds[i]})
	}

	return global.GormDB.Transaction(func(tx *gorm.DB) error {

		association := tx.Model(&model.SysUser{Id: req.Id}).Association("Roles")

		// 先需要清空用户关联的旧角色数据
		if err := association.Clear(); err != nil {
			return err
		}
		// 添加用户关联的角色
		if err := association.Append(roles); err != nil {
			return err
		}

		// 删除casbin用户角色
		if err := CasbinService.ClearRolesForUser(req.Id); err != nil {
			return err
		}
		// 添加casbin用户角色
		if err := CasbinService.AddRolesForUser(req.Id, req.RoleIds); err != nil {
			return err
		}

		return nil
	})

}
