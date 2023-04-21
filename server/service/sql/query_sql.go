package sql

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sqlReq "github.com/flipped-aurora/gin-vue-admin/server/model/sql/request"
    "gorm.io/gorm"
)

type QuerySqlService struct {
}

// CreateQuerySql 创建QuerySql记录
// Author [piexlmax](https://github.com/piexlmax)
func (querySqlService *QuerySqlService) CreateQuerySql(querySql *sql.QuerySql) (err error) {
	err = global.GVA_DB.Create(querySql).Error
	return err
}

// DeleteQuerySql 删除QuerySql记录
// Author [piexlmax](https://github.com/piexlmax)
func (querySqlService *QuerySqlService)DeleteQuerySql(querySql sql.QuerySql) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&sql.QuerySql{}).Where("id = ?", querySql.ID).Update("deleted_by", querySql.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&querySql).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteQuerySqlByIds 批量删除QuerySql记录
// Author [piexlmax](https://github.com/piexlmax)
func (querySqlService *QuerySqlService)DeleteQuerySqlByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&sql.QuerySql{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&sql.QuerySql{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateQuerySql 更新QuerySql记录
// Author [piexlmax](https://github.com/piexlmax)
func (querySqlService *QuerySqlService)UpdateQuerySql(querySql sql.QuerySql) (err error) {
	err = global.GVA_DB.Save(&querySql).Error
	return err
}

// GetQuerySql 根据id获取QuerySql记录
// Author [piexlmax](https://github.com/piexlmax)
func (querySqlService *QuerySqlService)GetQuerySql(id uint) (querySql sql.QuerySql, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&querySql).Error
	return
}

// GetQuerySqlInfoList 分页获取QuerySql记录
// Author [piexlmax](https://github.com/piexlmax)
func (querySqlService *QuerySqlService)GetQuerySqlInfoList(info sqlReq.QuerySqlSearch) (list []sql.QuerySql, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sql.QuerySql{})
    var querySqls []sql.QuerySql
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Cate != nil {
        db = db.Where("cate = ?",info.Cate)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&querySqls).Error
	return  querySqls, total, err
}
