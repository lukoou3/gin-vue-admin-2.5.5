package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
	"strings"
)

var AutoCodeSqlite = new(autoCodeSqlite)

type autoCodeSqlite struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeSqlite) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db = []response.Db{{Database: "db"}}
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeSqlite) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `SELECT name table_name FROM sqlite_master WHERE type='table' ORDER BY name`
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeSqlite) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column

	sql := fmt.Sprintf("select * from %s limit 0", tableName)

	var db *gorm.DB
	if businessDB == "" {
		db = global.GVA_DB
	} else {
		db = global.GVA_DBList[businessDB]
	}

	rows, _ := db.Raw(sql).Rows()
	defer rows.Close()

	columns, _ := rows.Columns()
	columnTypes, _ := rows.ColumnTypes()
	for i := range columns {
		col := columns[i]
		var tp string
		var tpLen string
		typeName := strings.ToLower(columnTypes[i].DatabaseTypeName())
		if strings.HasPrefix(typeName, "int") {
			tp = "int"
			tpLen = "20"
		} else if strings.HasPrefix(typeName, "varchar") || typeName == "text" {
			tp = "varchar"
			tpLen = "256"
		} else if typeName == "datetime" {
			tp = "datetime"
			tpLen = ""
		} else {
			return entities, fmt.Errorf("不支持的类型" + typeName)
		}
		entity := response.Column{DataType: tp, ColumnName: col, DataTypeLong: tpLen, ColumnComment: ""}
		entities = append(entities, entity)
	}
	return entities, nil
}
