package common

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

// BaseDO 基础模型
type BaseDO struct {
	UpdatedAt  time.Time      `gorm:"comment:修改时间" json:"-"`
	ModifyUser string         `gorm:"comment:修改人" json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"` // gorm逻辑删除
}

// BeforeUpdate 在更新之前执行
func (u *BaseDO) BeforeUpdate(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context

	if userId := USER_CTX.GetUserId(&ctx); userId != 0 {
		u.ModifyUser = strconv.FormatUint(userId, 10)
	}
	return nil
}

// BeforeSave 在保存之前执行
func (u *BaseDO) BeforeSave(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context

	if userId := USER_CTX.GetUserId(&ctx); userId != 0 {
		u.ModifyUser = strconv.FormatUint(userId, 10)
	}
	return nil
}
