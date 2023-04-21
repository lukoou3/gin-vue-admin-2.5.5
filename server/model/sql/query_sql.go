// 自动生成模板QuerySql
package sql

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// QuerySql 结构体
type QuerySql struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:128;"`
      Cate  *int `json:"cate" form:"cate" gorm:"column:cate;comment:分类;size:10;"`
      Sql  string `json:"sql" form:"sql" gorm:"column:sql;comment:sql;size:102400;"`
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:1024;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName QuerySql 表名
func (QuerySql) TableName() string {
  return "sql_querysql"
}

