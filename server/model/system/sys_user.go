package system

import (
	"gitee.com/nichanghao/gdmin/common"
)

type SysUser struct {
	Id       uint64 `gorm:"primarykey"`                            // 用户ID
	Username string `gorm:"index;type:varchar(128);comment:用户登录名"` // 用户登录名
	Password string `gorm:"type:varchar(64);comment:用户登录密码"`       // 用户登录密码
	Nickname string `gorm:"type:varchar(128);comment:用户昵称"`        // 用户昵称
	Phone    string `gorm:"type:varchar(16);comment:联系电话"`         // 联系电话
	Email    string `gorm:"type:varchar(64);comment:邮箱"`           // 邮箱
	common.BaseDO
}
