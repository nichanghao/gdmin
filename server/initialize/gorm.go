package initialize

import (
	"gitee.com/nichanghao/gdmin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func InitGorm() *gorm.DB {

	switch global.Config.Database.Driver {
	case "mysql":
		return initGormMysql()
	default:
		return initGormMysql()
	}
}

func initGormMysql() *gorm.DB {
	m := global.Config.Database.Mysql

	mysqlConfig := mysql.Config{
		DSN:                       m.DSN,
		DefaultStringSize:         256,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: m.TablePrefix,
		// 单数表名
		SingularTable: m.SingularTable,
	}}); err != nil {
		log.Fatalf("failed to connect database, the error is %v", err)
		return nil
	} else {
		s, _ := db.DB()
		s.SetMaxIdleConns(m.MaxIdleCount)
		s.SetMaxOpenConns(m.MaxOpenConns)

		return db
	}

}
