package common

import (
	"gorm.io/gorm"
	"time"
)

// BaseDO 基础模型
type BaseDO struct {
	UpdatedAt  time.Time      `gorm:"comment:修改时间"`
	ModifyUser string         `gorm:"comment:修改人"`
	DeletedAt  gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"` // gorm逻辑删除
}
