package system

import (
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
)

type SysMenuService struct {
}

func (service *SysMenuService) GetMenuTreeByUserId(userId uint64) ([]*model.SysMenu, error) {

	// 获取用户的菜单权限
	menuIds, err := CasbinService.GetPermissionMenuIdsByUserId(userId)
	if err != nil {
		return nil, err
	}

	var menus []*model.SysMenu

	global.GormDB.Where("id in (?)", menuIds).Find(&menus)
	if len(menus) == 0 {
		return nil, nil
	}

	return service.buildMenuTree(menus), nil

}

// 构建菜单树
func (*SysMenuService) buildMenuTree(menus []*model.SysMenu) []*model.SysMenu {
	// 创建一个 map 来存储每个菜单项
	menuMap := make(map[uint64]*model.SysMenu)
	for i := range menus {
		menuMap[menus[i].Id] = menus[i]
	}

	// 创建一个根菜单列表
	var rootMenus []*model.SysMenu

	for _, menu := range menus {
		if menu.ParentId == 0 {
			// 如果 ParentID 为 0，表示这是根菜单
			rootMenus = append(rootMenus, menu)
		} else {
			// 否则，找到父菜单并将其添加到父菜单的 Children 列表中
			if parentMenu, ok := menuMap[menu.ParentId]; ok {
				parentMenu.Children = append(parentMenu.Children, menu)
			}
		}
	}

	return rootMenus
}
