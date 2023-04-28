package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	//"github.com/glebarez/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GormSqlite() *gorm.DB {
	m := global.GVA_CONFIG.Sqlite
	if m.Path == "" {
		return nil
	}
	dbpath := m.Path

	if db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{}); err != nil {
		return nil
	} else {
		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormSqliteByConfig(m config.Sqlite) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	if m.Dbname == "" {
		return nil
	}
	dbpath := m.Dbname
	if len(m.Path) > 0 {
		dbpath = fmt.Sprintf("%s/%s", m.Path, m.Dbname)
	}

	if db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{}); err != nil {
		return nil
	} else {
		return db
	}
}
