package common

import (
	"gorm.io/gorm"
	"time"
)

type BaseDO struct {
	ModifyTime time.Time      `gorm:"comment:修改时间"`
	ModifyUser string         `gorm:"comment:修改人"`
	DeletedAt  gorm.DeletedAt `gorm:"index;comment:删除时间"` // gorm逻辑删除
}
