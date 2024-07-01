package system

import (
	"gitee.com/nichanghao/gdmin/model/common"
)

type SysUser struct {
	Id       uint64 `gorm:"primarykey" json:"id"`                                  // 用户ID
	Username string `gorm:"index;type:varchar(128);comment:用户登录名" json:"username"` // 用户登录名
	Password string `gorm:"type:varchar(64);comment:用户登录密码" json:"-"`              // 用户登录密码
	Nickname string `gorm:"type:varchar(128);comment:用户昵称" json:"nickname"`        // 用户昵称
	Phone    string `gorm:"type:varchar(16);comment:联系电话" json:"phone"`            // 联系电话
	Email    string `gorm:"type:varchar(64);comment:邮箱" json:"email"`              // 邮箱
	common.BaseDO
}
