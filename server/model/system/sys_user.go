package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysUser struct {
	Id       uint64    `gorm:"primarykey;comment:用户ID" json:"id"`
	Username string    `gorm:"index;type:varchar(128);comment:用户登录名" json:"username"`
	Password string    `gorm:"type:varchar(64);comment:用户登录密码" json:"-"`
	Nickname string    `gorm:"type:varchar(128);comment:用户昵称" json:"nickname"`
	Phone    string    `gorm:"type:varchar(16);comment:联系电话" json:"phone"`
	Email    string    `gorm:"type:varchar(64);comment:邮箱" json:"email"`
	Roles    []SysRole `gorm:"many2many:sys_user_role;" json:"roles"` // 用户角色关系
	common.BaseDO
}
