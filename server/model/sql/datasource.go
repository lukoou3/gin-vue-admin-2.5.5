// 自动生成模板Datasource
package sql

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Datasource 结构体
type Datasource struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:数据源名称;size:128;"`
      Alias  string `json:"alias" form:"alias" gorm:"column:alias;comment:数据源别名;size:128;"`
      Cate  *int `json:"cate" form:"cate" gorm:"column:cate;comment:分类;"`
      Introduction  string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:简介;size:2048;"`
      Sql  string `json:"sql" form:"sql" gorm:"column:sql;comment:sql;size:10240;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Datasource 表名
func (Datasource) TableName() string {
  return "sql_datasource"
}

