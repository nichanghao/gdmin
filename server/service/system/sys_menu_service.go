package system

import (
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gorm.io/gorm"
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

func (*SysMenuService) AddMenu(menu *model.SysMenu) (uint64, error) {

	if err := global.GormDB.Create(menu).Error; err != nil {
		return 0, err
	}

	return menu.Id, nil
}

func (*SysMenuService) EditMenu(menu *model.SysMenu) (uint64, error) {

	if err := global.GormDB.Where("id =?", menu.Id).Updates(menu).Error; err != nil {
		return 0, err
	}

	return menu.Id, nil
}

func (*SysMenuService) DeleteMenu(menuId uint64) error {

	var count int64
	if err := global.GormDB.Model(&model.SysMenu{}).Where("parent_id = ?", menuId).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return buserr.NewNoticeBusErr("菜单下有子菜单，不能被删除")
	}

	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		// 删除casbin权限
		if err := CasbinService.DeletePermissionByMenuId(menuId); err != nil {
			return err
		}

		// 删除菜单
		if err := global.GormDB.Delete(&model.SysMenu{}, menuId).Error; err != nil {
			return err
		}
		return nil
	})

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
