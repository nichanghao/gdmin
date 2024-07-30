package system

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/utils"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type SysUserService struct {
}

// Login 用户登录
func (userService *SysUserService) Login(user *model.SysUser) (*response.SysUserLoginResp, error) {

	var userRes *model.SysUser

	if err := global.GormDB.Where("username = ?", user.Username).First(&userRes).Error; err != nil {
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

// GetSelfUserInfo 获取当前用户信息
func (userService *SysUserService) GetSelfUserInfo(id uint64) (res *model.SysUser, err error) {

	if err = global.GormDB.Model(&model.SysUser{Id: id}).Preload("Roles").First(&res).Error; err != nil {
		return
	}

	return
}

// PageUsers 分页查询用户列表
func (userService *SysUserService) PageUsers(req *request.SysUserPageReq) (*common.PageResp, error) {
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
	if req.Status != 0 {
		tx.Where("status = ?", req.Status)
	}
	if req.Gender != 0 {
		tx.Where("gender = ?", req.Gender)
	}

	res := &common.PageResp{Current: req.Current, Size: req.Size, Records: make([]any, 0)}
	// 查询数量
	if err := tx.Count(&res.Total).Error; err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}

	var userList []*model.SysUser
	if err := tx.Preload("Roles").Find(&userList).Error; err != nil {
		return res, err
	} else {
		res.Records = userList
	}

	return res, nil
}

// AddUser 新增用户
func (userService *SysUserService) AddUser(req *common.Request) error {
	var user model.SysUser
	if err := copier.Copy(&user, req.Data); err != nil {
		return err
	}

	// 加密密码
	password, err := utils.BCRYPT.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password

	return global.GormDB.Create(&user).Error
}

// EditUser 编辑用户
func (userService *SysUserService) EditUser(req *common.Request) error {

	var user model.SysUser
	if err := copier.Copy(&user, req.Data); err != nil {
		return err
	}

	if err := global.GormDB.Model(&model.SysUser{}).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

// ResetPassword 重置密码
func (userService *SysUserService) ResetPassword(req *common.Request) error {
	resetPwd := req.Data.(*request.SysUserResetPwdReq)

	password, err := utils.BCRYPT.HashPassword(resetPwd.Password)
	if err != nil {
		return err
	}

	if err = global.GormDB.WithContext(req.Context).Model(&model.SysUser{}).Where("id = ?", resetPwd.Id).Update("password", password).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser 删除用户
func (userService *SysUserService) DeleteUser(req *common.Request) error {
	userId := req.Data.(*request.QueryIdReq).Id

	return global.GormDB.Transaction(func(tx *gorm.DB) error {

		// 删除用户并同步删除关联的角色
		if err := tx.WithContext(req.Context).Select("Roles").Delete(&model.SysUser{Id: userId}).Error; err != nil {
			return err
		}

		// 删除casbin用户
		if err := CasbinService.ClearRolesForUser(userId); err != nil {
			return err
		}

		return nil
	})
}

// AssignRoles 分配角色给用户
func (userService *SysUserService) AssignRoles(req *common.Request) error {
	assignRole := req.Data.(*request.SysUserAssignRoleReq)

	var roles []*model.SysRole
	for i := range assignRole.RoleIds {
		roles = append(roles, &model.SysRole{Id: assignRole.RoleIds[i]})
	}

	return global.GormDB.Transaction(func(tx *gorm.DB) error {

		// 清空用户关联的旧角色数据
		if err := tx.Model(&model.SysUser{Id: assignRole.Id}).Association("Roles").Clear(); err != nil {
			return err
		}
		// 删除casbin用户角色
		if err := CasbinService.ClearRolesForUser(assignRole.Id); err != nil {
			return err
		}

		// 移除关联的所有角色时，直接返回
		if len(roles) == 0 {
			return nil
		}
		// 添加用户关联的角色
		if err := tx.Model(&model.SysUser{Id: assignRole.Id}).Association("Roles").Append(roles); err != nil {
			return err
		}
		// 添加casbin用户角色
		if err := CasbinService.AddRolesForUser(assignRole.Id, assignRole.RoleIds); err != nil {
			return err
		}

		return nil
	})

}

// UpdateStatus 更新用户状态
func (userService *SysUserService) UpdateStatus(req *request.SysUserEditReq) error {

	if err := global.GormDB.Model(&model.SysUser{}).Where("id = ?", req.Id).Update("status", req.Status).Error; err != nil {
		return err
	}

	return nil
}
