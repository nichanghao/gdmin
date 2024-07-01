package system

import "gitee.com/nichanghao/gdmin/model/common"

type SysRole struct {
	Id   uint64 `gorm:"primarykey;comment:角色ID"`       // 角色ID
	Name string `gorm:"type:varchar(32);comment:角色名"`  // 角色名
	Code string `gorm:"type:varchar(32);comment:角色标识"` // 角色标识
	common.BaseDO
}
