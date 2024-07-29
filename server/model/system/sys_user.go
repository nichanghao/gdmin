package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysUser struct {
	Id       uint64    `gorm:"primarykey;comment:用户ID" json:"id"`
	Username string    `gorm:"index;type:varchar(128);comment:用户登录名" json:"username"`
	Password string    `gorm:"type:varchar(64);comment:用户登录密码" json:"-"`
	Nickname string    `gorm:"type:varchar(128);comment:用户昵称" json:"nickname"`
	Gender   uint8     `gorm:"type:tinyint(1);comment:性别(1:男,2:女)" json:"gender"`
	Phone    string    `gorm:"type:varchar(16);comment:联系电话" json:"phone"`
	Email    string    `gorm:"type:varchar(64);comment:邮箱" json:"email"`
	Status   uint8     `gorm:"type:tinyint(1);default:1;comment:用户状态(1:正常,2:停用)" json:"status"`
	Roles    []SysRole `gorm:"many2many:sys_user_role;" json:"roles"` // 用户角色关系
	common.BaseDO
}
