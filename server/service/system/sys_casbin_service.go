package system

import (
	"gitee.com/nichanghao/gdmin/global"
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
func (casbinService *SysCasbinService) GetPermissionMenuIdsByUserId(userId uint64) ([]string, error) {

	policies, err := global.Enforcer.GetImplicitPermissionsForUser(casbinService.GetCasbinUserStr(userId))
	menuIds := make([]string, 0, len(policies))

	if err != nil {
		return menuIds, err
	}
	if len(policies) == 0 {
		return menuIds, nil
	}

	for i := range policies {
		menuIds = append(menuIds, policies[i][2])
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

func (*SysCasbinService) GetCasbinUserStr(userId uint64) string {
	return userPrefix + strconv.FormatUint(userId, 10)
}

func (*SysCasbinService) GetCasbinRoleStr(roleId uint64) string {
	return rolePrefix + strconv.FormatUint(roleId, 10)
}
