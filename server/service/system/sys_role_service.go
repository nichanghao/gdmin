package system

import (
	"errors"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/model/common"
	"gitee.com/nichanghao/gdmin/web/request"
	"gorm.io/gorm"
)

type SysRoleService struct {
}

// PageRoles 分页查询角色列表
func (*SysRoleService) PageRoles(req *request.SysRolePageReq) ([]*model.SysRole, error) {

	tx := global.GormDB.Model(&model.SysRole{}).Limit(req.Limit).Offset(req.Offset)
	if req.Name != "" {
		tx.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Code != "" {
		tx.Where("code LIKE ?", "%"+req.Code+"%")
	}
	var roleList []*model.SysRole

	if err := tx.Find(&roleList).Error; err != nil {
		return roleList, err
	}

	return roleList, nil

}

// AddRole 新增角色
func (roleService *SysRoleService) AddRole(role *model.SysRole) error {

	return global.GormDB.Model(&model.SysRole{}).Transaction(func(tx *gorm.DB) error {

		if err := roleService.validateDuplicateRole(tx, role); err != nil {
			return err
		}

		if err := tx.Create(role).Error; err != nil {
			return err
		}

		return nil
	})
}

// EditRole 编辑角色
func (roleService *SysRoleService) EditRole(role *model.SysRole) error {
	return global.GormDB.Model(&model.SysRole{}).Transaction(func(tx *gorm.DB) error {

		var roleOld = model.SysRole{}
		if errors.Is(tx.Where("id = ?", role.Id).First(&roleOld).Error, gorm.ErrRecordNotFound) {
			return common.NewNoticeBusErr("该角色不存在！")
		}

		if roleOld.Name != role.Name {
			if err := roleService.validateDuplicateRoleByName(tx, role.Name); err != nil {
				return err
			}
		}

		if roleOld.Code != role.Code {
			if err := roleService.validateDuplicateRoleByCode(tx, role.Code); err != nil {
				return err
			}
		}

		if err := tx.Where("id = ?", role.Id).Updates(role).Error; err != nil {
			return err
		}

		return nil
	})
}

// DeleteRole 删除角色
func (roleService *SysRoleService) DeleteRole(roleId uint64) error {

	return global.GormDB.Model(&model.SysRole{}).Transaction(func(tx *gorm.DB) error {
		var role = model.SysRole{}

		if errors.Is(tx.Where("id = ?", roleId).Preload("Users").First(&role).Error, gorm.ErrRecordNotFound) {
			return common.NewNoticeBusErr("该角色不存在！")
		}
		if len(role.Users) > 0 {
			return common.NewNoticeBusErr("该角色已分配给用户，不能删除！")
		}

		return tx.Delete(&model.SysRole{}, roleId).Error
	})

}

// 校验角色名称和编码是否重复
func (roleService *SysRoleService) validateDuplicateRole(tx *gorm.DB, role *model.SysRole) error {
	if err := roleService.validateDuplicateRoleByName(tx, role.Name); err != nil {
		return err
	}

	if err := roleService.validateDuplicateRoleByCode(tx, role.Code); err != nil {
		return err
	}

	return nil
}

func (*SysRoleService) validateDuplicateRoleByName(tx *gorm.DB, name string) error {
	var count int64

	if err := tx.Where("name = ?", name).Limit(1).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return common.NewNoticeBusErr("角色名称已存在！")
	}

	return nil
}

func (*SysRoleService) validateDuplicateRoleByCode(tx *gorm.DB, code string) error {
	var count int64

	if err := tx.Where("code = ?", code).Limit(1).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return common.NewNoticeBusErr("角色编码已存在！")
	}

	return nil
}
