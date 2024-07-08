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
	role_prefix = "r:"

	userPrefix = "u:"
)

type SysCasbinService struct{}

// GetPermissionMenuIdsByUserId 获取用户菜单权限
func (casbinService *SysCasbinService) GetPermissionMenuIdsByUserId(userId uint64) ([]string, error) {

	policies, err := global.Enforcer.GetImplicitPermissionsForUser(casbinService.getCasbinUserStr(userId))
	menuIds := make([]string, 0, len(policies))

	if err != nil {
		return menuIds, err
	}
	if len(policies) == 0 {
		return menuIds, nil
	}

	for i := range policies {
		menuIds = append(menuIds, policies[i][3])
	}

	return menuIds, nil

}

func (casbinService *SysCasbinService) DeletePermissionByMenuId(menuId uint64) error {

	_, err := global.Enforcer.RemoveFilteredPolicy(3, strconv.FormatUint(menuId, 10))
	if err != nil {
		return err
	}

	return nil
}

func (*SysCasbinService) getCasbinUserStr(userId uint64) string {
	return userPrefix + strconv.FormatUint(userId, 10)
}