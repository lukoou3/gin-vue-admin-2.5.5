// 自动生成模板Shellcode
package code

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Shellcode 结构体
type Shellcode struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:128;"`
      Code  string `json:"code" form:"code" gorm:"column:code;comment:code;size:102400;"`
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:10240;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Shellcode 表名
func (Shellcode) TableName() string {
  return "code_shellcode"
}

