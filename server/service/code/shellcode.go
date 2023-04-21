package code

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/code"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    codeReq "github.com/flipped-aurora/gin-vue-admin/server/model/code/request"
    "gorm.io/gorm"
)

type ShellcodeService struct {
}

// CreateShellcode 创建Shellcode记录
// Author [piexlmax](https://github.com/piexlmax)
func (shellcodeService *ShellcodeService) CreateShellcode(shellcode *code.Shellcode) (err error) {
	err = global.GVA_DB.Create(shellcode).Error
	return err
}

// DeleteShellcode 删除Shellcode记录
// Author [piexlmax](https://github.com/piexlmax)
func (shellcodeService *ShellcodeService)DeleteShellcode(shellcode code.Shellcode) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&code.Shellcode{}).Where("id = ?", shellcode.ID).Update("deleted_by", shellcode.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&shellcode).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteShellcodeByIds 批量删除Shellcode记录
// Author [piexlmax](https://github.com/piexlmax)
func (shellcodeService *ShellcodeService)DeleteShellcodeByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&code.Shellcode{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&code.Shellcode{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateShellcode 更新Shellcode记录
// Author [piexlmax](https://github.com/piexlmax)
func (shellcodeService *ShellcodeService)UpdateShellcode(shellcode code.Shellcode) (err error) {
	err = global.GVA_DB.Save(&shellcode).Error
	return err
}

// GetShellcode 根据id获取Shellcode记录
// Author [piexlmax](https://github.com/piexlmax)
func (shellcodeService *ShellcodeService)GetShellcode(id uint) (shellcode code.Shellcode, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shellcode).Error
	return
}

// GetShellcodeInfoList 分页获取Shellcode记录
// Author [piexlmax](https://github.com/piexlmax)
func (shellcodeService *ShellcodeService)GetShellcodeInfoList(info codeReq.ShellcodeSearch) (list []code.Shellcode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&code.Shellcode{})
    var shellcodes []code.Shellcode
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&shellcodes).Error
	return  shellcodes, total, err
}
