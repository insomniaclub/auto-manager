package initializer

import (
	"auto-manager/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitGorm() *gorm.DB {
	m := global.AM_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func gormConfig(mod bool) *gorm.Config {
	switch global.AM_CONFIG.Mysql.LogZap {
	case "Silent":
		return nil
	case "Error":
		return nil
	case "Warn":
		return nil
	case "Info":
		return nil
	default:
		return nil
	}
}