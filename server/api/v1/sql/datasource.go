package sql

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/sql"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sqlReq "github.com/flipped-aurora/gin-vue-admin/server/model/sql/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type DatasourceApi struct {
}

var datasourceService = service.ServiceGroupApp.SqlServiceGroup.DatasourceService


// CreateDatasource 创建Datasource
// @Tags Datasource
// @Summary 创建Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sql.Datasource true "创建Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /datasource/createDatasource [post]
func (datasourceApi *DatasourceApi) CreateDatasource(c *gin.Context) {
	var datasource sql.Datasource
	err := c.ShouldBindJSON(&datasource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    datasource.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Cate":{utils.NotEmpty()},
    }
	if err := utils.Verify(datasource, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := datasourceService.CreateDatasource(&datasource); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDatasource 删除Datasource
// @Tags Datasource
// @Summary 删除Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sql.Datasource true "删除Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /datasource/deleteDatasource [delete]
func (datasourceApi *DatasourceApi) DeleteDatasource(c *gin.Context) {
	var datasource sql.Datasource
	err := c.ShouldBindJSON(&datasource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    datasource.DeletedBy = utils.GetUserID(c)
	if err := datasourceService.DeleteDatasource(datasource); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDatasourceByIds 批量删除Datasource
// @Tags Datasource
// @Summary 批量删除Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /datasource/deleteDatasourceByIds [delete]
func (datasourceApi *DatasourceApi) DeleteDatasourceByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := datasourceService.DeleteDatasourceByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDatasource 更新Datasource
// @Tags Datasource
// @Summary 更新Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sql.Datasource true "更新Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /datasource/updateDatasource [put]
func (datasourceApi *DatasourceApi) UpdateDatasource(c *gin.Context) {
	var datasource sql.Datasource
	err := c.ShouldBindJSON(&datasource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    datasource.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Cate":{utils.NotEmpty()},
      }
    if err := utils.Verify(datasource, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := datasourceService.UpdateDatasource(datasource); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDatasource 用id查询Datasource
// @Tags Datasource
// @Summary 用id查询Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sql.Datasource true "用id查询Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /datasource/findDatasource [get]
func (datasourceApi *DatasourceApi) FindDatasource(c *gin.Context) {
	var datasource sql.Datasource
	err := c.ShouldBindQuery(&datasource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redatasource, err := datasourceService.GetDatasource(datasource.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redatasource": redatasource}, c)
	}
}

// GetDatasourceList 分页获取Datasource列表
// @Tags Datasource
// @Summary 分页获取Datasource列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sqlReq.DatasourceSearch true "分页获取Datasource列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /datasource/getDatasourceList [get]
func (datasourceApi *DatasourceApi) GetDatasourceList(c *gin.Context) {
	var pageInfo sqlReq.DatasourceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := datasourceService.GetDatasourceInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
