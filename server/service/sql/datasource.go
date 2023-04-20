package sql

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sqlReq "github.com/flipped-aurora/gin-vue-admin/server/model/sql/request"
    "gorm.io/gorm"
)

type DatasourceService struct {
}

// CreateDatasource 创建Datasource记录
// Author [piexlmax](https://github.com/piexlmax)
func (datasourceService *DatasourceService) CreateDatasource(datasource *sql.Datasource) (err error) {
	err = global.GVA_DB.Create(datasource).Error
	return err
}

// DeleteDatasource 删除Datasource记录
// Author [piexlmax](https://github.com/piexlmax)
func (datasourceService *DatasourceService)DeleteDatasource(datasource sql.Datasource) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&sql.Datasource{}).Where("id = ?", datasource.ID).Update("deleted_by", datasource.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&datasource).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteDatasourceByIds 批量删除Datasource记录
// Author [piexlmax](https://github.com/piexlmax)
func (datasourceService *DatasourceService)DeleteDatasourceByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&sql.Datasource{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&sql.Datasource{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateDatasource 更新Datasource记录
// Author [piexlmax](https://github.com/piexlmax)
func (datasourceService *DatasourceService)UpdateDatasource(datasource sql.Datasource) (err error) {
	err = global.GVA_DB.Save(&datasource).Error
	return err
}

// GetDatasource 根据id获取Datasource记录
// Author [piexlmax](https://github.com/piexlmax)
func (datasourceService *DatasourceService)GetDatasource(id uint) (datasource sql.Datasource, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&datasource).Error
	return
}

// GetDatasourceInfoList 分页获取Datasource记录
// Author [piexlmax](https://github.com/piexlmax)
func (datasourceService *DatasourceService)GetDatasourceInfoList(info sqlReq.DatasourceSearch) (list []sql.Datasource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sql.Datasource{})
    var datasources []sql.Datasource
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Alias != "" {
        db = db.Where("alias LIKE ?","%"+ info.Alias+"%")
    }
    if info.Cate != nil {
        db = db.Where("cate = ?",info.Cate)
    }
    if info.Introduction != "" {
        db = db.Where("introduction LIKE ?","%"+ info.Introduction+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&datasources).Error
	return  datasources, total, err
}
