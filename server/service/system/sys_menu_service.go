package system

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model"
	"gitee.com/nichanghao/gdmin/web/request"
	"gitee.com/nichanghao/gdmin/web/response"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	MenuService = new(SysMenuService)
)

type SysMenuService struct {
}

// GetAllMenuTree 获取所有菜单树
func (service *SysMenuService) GetAllMenuTree() (res []*model.SysMenu, err error) {

	var menus []*model.SysMenu
	if err = global.GormDB.Find(&menus).Error; err != nil {
		return
	}

	res = service.buildMenuTree(menus)
	return
}

func (*SysMenuService) AddMenu(req *common.Request) (uint64, error) {

	var menu model.SysMenu
	err := copier.Copy(&menu, req.Data)
	if err != nil {
		return 0, err
	}

	if err = global.GormDB.WithContext(req.Context).Create(&menu).Error; err != nil {
		return 0, err
	}

	return menu.Id, nil
}

func (*SysMenuService) EditMenu(req *common.Request) (uint64, error) {

	var menu model.SysMenu
	err := copier.Copy(&menu, req.Data)
	if err != nil {
		return 0, err
	}

	if err = global.GormDB.WithContext(req.Context).Where("id =?", menu.Id).Updates(&menu).Error; err != nil {
		return 0, err
	}

	return menu.Id, nil
}

func (*SysMenuService) DeleteMenu(req *common.Request) error {

	menuId := req.Data.(*request.SysMenuUpdateReq).Id

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
		if err := tx.WithContext(req.Context).Delete(&model.SysMenu{}, menuId).Error; err != nil {
			return err
		}
		return nil
	})

}

// GetSelfPermissionRouters 获取自身权限路由
func (service *SysMenuService) GetSelfPermissionRouters(userId uint64) (res *response.SysPermissionRoutersResp, err error) {
	res = &response.SysPermissionRoutersResp{Home: "home"}

	// 获取用户的菜单权限
	menuIds, err := CasbinService.GetPermissionMenuIdsByUserId(userId)
	if err != nil {
		return
	}

	var routes []*response.SysRoutesResp
	tx := global.GormDB.Model(&model.SysMenu{}).Where("status = 1 AND id IN (?)", menuIds)
	tx.Select("id, route_name as name, parent_id, path, component, permission, type, meta").Find(&routes)
	if len(routes) == 0 {
		return
	}

	res.Routes, res.Permissions = service.buildPermissionRoutes(routes)
	return
}

// buildPermissionRoutes 构建权限路由
func (*SysMenuService) buildPermissionRoutes(routes []*response.SysRoutesResp) ([]*response.SysRoutesResp, []string) {
	// 创建一个 map 来存储每个菜单项
	routeMap := make(map[uint64]*response.SysRoutesResp)
	for i := range routes {
		routeMap[routes[i].Id] = routes[i]
	}

	// 创建一个根菜单列表
	var rootRoutes = make([]*response.SysRoutesResp, 0, len(routes))

	// 权限列表
	var permissionList = make([]string, 0, len(routes))

	for _, menu := range routes {

		if menu.Permission != "" {
			permissionList = append(permissionList, menu.Permission)
		}

		if menu.Type == 3 {
			continue
		}

		if menu.ParentId == 0 {
			// 如果 ParentID 为 0，表示这是根菜单
			rootRoutes = append(rootRoutes, menu)
		} else {
			// 否则，找到父菜单并将其添加到父菜单的 Children 列表中
			if parentMenu, ok := routeMap[menu.ParentId]; ok {
				parentMenu.Children = append(parentMenu.Children, menu)
			}
		}
	}

	return rootRoutes, permissionList
}

// 构建菜单树
func (*SysMenuService) buildMenuTree(menus []*model.SysMenu) []*model.SysMenu {
	// 创建一个 map 来存储每个菜单项
	menuMap := make(map[uint64]*model.SysMenu)
	for i := range menus {
		menuMap[menus[i].Id] = menus[i]
	}

	// 创建一个根菜单列表
	var rootMenus = make([]*model.SysMenu, 0, len(menus))

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
