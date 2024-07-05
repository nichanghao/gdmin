package system

import "gitee.com/nichanghao/gdmin/model/common"

type SysRole struct {
	Id    uint64    `gorm:"primarykey;comment:角色ID"`
	Name  string    `gorm:"type:varchar(32);comment:角色名"`
	Code  string    `gorm:"type:varchar(32);comment:角色标识"`
	Desc  string    `gorm:"type:varchar(255);comment:备注"`
	Users []SysUser `gorm:"many2many:sys_user_role;"` // 角色与用户的多对多关系
	common.BaseDO
}
