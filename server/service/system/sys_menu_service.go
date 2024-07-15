package system

import (
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gorm.io/gorm"
	"sort"
)

var (
	MenuService = new(SysMenuService)
)

type SysMenuService struct {
}

// GetAllMenuTree 获取所有菜单树
func (service *SysMenuService) GetAllMenuTree() (res []*model.SysMenu, err error) {

	if err = global.GormDB.Find(&res).Error; err != nil {
		return
	}

	tree, _ := service.buildMenuTree(res, false)
	return tree, nil
}

// GetMenuTreeByUserId 获取用户的菜单树
func (service *SysMenuService) GetMenuTreeByUserId(userId uint64) ([]*model.SysMenu, []string, error) {

	// 获取用户的菜单权限
	menuIds, err := CasbinService.GetPermissionMenuIdsByUserId(userId)
	if err != nil {
		return nil, nil, err
	}

	var menus []*model.SysMenu

	global.GormDB.Where("id in (?)", menuIds).Find(&menus)
	if len(menus) == 0 {
		return nil, nil, nil
	}

	tree, permissionList := service.buildMenuTree(menus, true)
	return tree, permissionList, nil

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
func (*SysMenuService) buildMenuTree(menus []*model.SysMenu, excludeBtn bool) ([]*model.SysMenu, []string) {
	// 创建一个 map 来存储每个菜单项
	menuMap := make(map[uint64]*model.SysMenu)
	for i := range menus {
		menuMap[menus[i].Id] = menus[i]
	}

	// 排序
	sort.Slice(menus, func(i, j int) bool { return menus[i].Sort < menus[j].Sort })

	// 创建一个根菜单列表
	var rootMenus = make([]*model.SysMenu, 0, len(menus))

	// 权限列表
	var permissionList = make([]string, 0, len(menus))

	for _, menu := range menus {

		if menu.Permission != "" {
			permissionList = append(permissionList, menu.Permission)
		}

		// 是否排除按钮
		if excludeBtn && menu.Type == 2 {
			continue
		}

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

	return rootMenus, permissionList
}
