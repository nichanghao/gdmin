package system

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"strconv"
)

var (
	CasbinService = new(SysCasbinService)
)

// casbin 中 区分角色和用户的前缀
const (
	rolePrefix = "r:"

	userPrefix = "u:"
)

type SysCasbinService struct{}

// GetPermissionMenuIdsByUserId 获取用户菜单权限
func (casbinService *SysCasbinService) GetPermissionMenuIdsByUserId(userId uint64) ([]uint64, error) {

	policies, err := global.Enforcer.GetImplicitPermissionsForUser(casbinService.GetCasbinUserStr(userId))
	menuIds := make([]uint64, 0, len(policies))

	if err != nil {
		return menuIds, err
	}
	if len(policies) == 0 {
		return menuIds, nil
	}

	for i := range policies {
		num, _ := strconv.ParseUint(policies[i][2], 10, 64)
		menuIds = append(menuIds, num)
	}

	return menuIds, nil

}

// GetPermissionMenuIdsByRoleId 获取角色菜单权限
func (casbinService *SysCasbinService) GetPermissionMenuIdsByRoleId(roleId uint64) ([]uint64, error) {

	policies, err := global.Enforcer.GetFilteredPolicy(0, casbinService.GetCasbinRoleStr(roleId))
	menuIds := make([]uint64, 0, len(policies))

	if err != nil {
		return menuIds, err
	}
	if len(policies) == 0 {
		return menuIds, nil
	}

	for i := range policies {
		num, _ := strconv.ParseUint(policies[i][2], 10, 64)
		menuIds = append(menuIds, num)
	}

	return menuIds, nil

}

// DeletePermissionByMenuId 删除菜单权限
func (casbinService *SysCasbinService) DeletePermissionByMenuId(menuId uint64) error {

	_, err := global.Enforcer.RemoveFilteredPolicy(2, strconv.FormatUint(menuId, 10))
	if err != nil {
		return err
	}

	return nil
}

// ClearRolesForUser 删除casbin用户所有角色
func (casbinService *SysCasbinService) ClearRolesForUser(userId uint64) error {

	_, err := global.Enforcer.DeleteUser(casbinService.GetCasbinUserStr(userId))
	if err != nil {
		return err
	}

	return nil
}

// AddRolesForUser casbin用户添加角色
func (casbinService *SysCasbinService) AddRolesForUser(userId uint64, roleIds []uint64) error {

	roleStrs := make([]string, 0, len(roleIds))
	for i := range roleIds {
		roleStrs = append(roleStrs, casbinService.GetCasbinRoleStr(roleIds[i]))
	}

	if _, err := global.Enforcer.AddRolesForUser(casbinService.GetCasbinUserStr(userId), roleStrs); err != nil {
		return err
	}

	return nil
}

// GetPermissionByUserId 获取用户所有权限
func (casbinService *SysCasbinService) GetPermissionByUserId(userId uint64) (map[string]any, error) {

	policies, err := global.Enforcer.GetImplicitPermissionsForUser(casbinService.GetCasbinUserStr(userId))

	permissionSet := make(map[string]any, len(policies))
	if err != nil {
		return permissionSet, err
	}

	for i := range policies {
		permissionSet[policies[i][1]] = struct{}{}
	}

	return permissionSet, nil

}

// DeletePermissionByRoleAndMenus 删除角色菜单权限
func (casbinService *SysCasbinService) DeletePermissionByRoleAndMenus(roleId uint64, menus []uint64) (err error) {
	roleIdStr := casbinService.GetCasbinRoleStr(roleId)
	for i := range menus {
		_, err = global.Enforcer.RemoveFilteredPolicy(0, roleIdStr, "", strconv.FormatUint(menus[i], 10))
	}
	return
}

// AddPermissionByRoleAndMenus 添加角色菜单权限
func (casbinService *SysCasbinService) AddPermissionByRoleAndMenus(roleId uint64, menus []model.SysMenu) (err error) {
	var policies [][]string
	casbinRoleStr := CasbinService.GetCasbinRoleStr(roleId)
	for i := range menus {
		policies[i] = []string{casbinRoleStr, menus[i].Permission, strconv.FormatUint(menus[i].Id, 10)}
	}
	_, err = global.Enforcer.AddPolicies(policies)

	return
}

func (*SysCasbinService) GetCasbinUserStr(userId uint64) string {
	return userPrefix + strconv.FormatUint(userId, 10)
}

func (*SysCasbinService) GetCasbinRoleStr(roleId uint64) string {
	return rolePrefix + strconv.FormatUint(roleId, 10)
}
