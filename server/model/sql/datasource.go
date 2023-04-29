// 自动生成模板Datasource
package sql

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	system "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Datasource 结构体
type Datasource struct {
	global.GVA_MODEL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:数据源名称;size:128;"`
	Alias        string `json:"alias" form:"alias" gorm:"column:alias;comment:数据源别名;size:128;"`
	Cate         *int   `json:"cate" form:"cate" gorm:"column:cate;comment:分类;"`
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:简介;size:2048;"`
	Sql          string `json:"sql" form:"sql" gorm:"column:sql;comment:sql;size:10240;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Datasource 表名
func (*Datasource) TableName() string {
	return "sql_datasource"
}

func (*Datasource) RegisterApis(db *gorm.DB) {
	abbreviation := "datasource" // Struct简称
	structName := "Datasource"   // Struct名称
	description := "数据源"         // Struct中文名称

	apiList := []system.SysApi{
		{
			Path:        "/" + abbreviation + "/" + "create" + structName,
			Description: "新增" + description,
			ApiGroup:    abbreviation,
			Method:      "POST",
		},
		{
			Path:        "/" + abbreviation + "/" + "delete" + structName,
			Description: "删除" + description,
			ApiGroup:    abbreviation,
			Method:      "DELETE",
		},
		{
			Path:        "/" + abbreviation + "/" + "delete" + structName + "ByIds",
			Description: "批量删除" + description,
			ApiGroup:    abbreviation,
			Method:      "DELETE",
		},
		{
			Path:        "/" + abbreviation + "/" + "update" + structName,
			Description: "更新" + description,
			ApiGroup:    abbreviation,
			Method:      "PUT",
		},
		{
			Path:        "/" + abbreviation + "/" + "find" + structName,
			Description: "根据ID获取" + description,
			ApiGroup:    abbreviation,
			Method:      "GET",
		},
		{
			Path:        "/" + abbreviation + "/" + "get" + structName + "List",
			Description: "获取" + description + "列表",
			ApiGroup:    abbreviation,
			Method:      "GET",
		},
	}
	for _, v := range apiList {
		var api system.SysApi
		if errors.Is(db.Where("path = ? AND method = ?", v.Path, v.Method).First(&api).Error, gorm.ErrRecordNotFound) {
			if err := db.Create(&v).Error; err != nil {
				global.GVA_LOG.Error("register api failed"+v.Path+" - "+v.Method, zap.Error(err))
			}

			global.GVA_LOG.Info("register api: " + v.Path + " - " + v.Method)
		}
	}
}
