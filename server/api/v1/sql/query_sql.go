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

type QuerySqlApi struct {
}

var querySqlService = service.ServiceGroupApp.SqlServiceGroup.QuerySqlService


// CreateQuerySql 创建QuerySql
// @Tags QuerySql
// @Summary 创建QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sql.QuerySql true "创建QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /querySql/createQuerySql [post]
func (querySqlApi *QuerySqlApi) CreateQuerySql(c *gin.Context) {
	var querySql sql.QuerySql
	err := c.ShouldBindJSON(&querySql)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    querySql.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Cate":{utils.NotEmpty()},
    }
	if err := utils.Verify(querySql, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := querySqlService.CreateQuerySql(&querySql); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuerySql 删除QuerySql
// @Tags QuerySql
// @Summary 删除QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sql.QuerySql true "删除QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /querySql/deleteQuerySql [delete]
func (querySqlApi *QuerySqlApi) DeleteQuerySql(c *gin.Context) {
	var querySql sql.QuerySql
	err := c.ShouldBindJSON(&querySql)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    querySql.DeletedBy = utils.GetUserID(c)
	if err := querySqlService.DeleteQuerySql(querySql); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuerySqlByIds 批量删除QuerySql
// @Tags QuerySql
// @Summary 批量删除QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /querySql/deleteQuerySqlByIds [delete]
func (querySqlApi *QuerySqlApi) DeleteQuerySqlByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := querySqlService.DeleteQuerySqlByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuerySql 更新QuerySql
// @Tags QuerySql
// @Summary 更新QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sql.QuerySql true "更新QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /querySql/updateQuerySql [put]
func (querySqlApi *QuerySqlApi) UpdateQuerySql(c *gin.Context) {
	var querySql sql.QuerySql
	err := c.ShouldBindJSON(&querySql)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    querySql.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Cate":{utils.NotEmpty()},
      }
    if err := utils.Verify(querySql, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := querySqlService.UpdateQuerySql(querySql); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuerySql 用id查询QuerySql
// @Tags QuerySql
// @Summary 用id查询QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sql.QuerySql true "用id查询QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /querySql/findQuerySql [get]
func (querySqlApi *QuerySqlApi) FindQuerySql(c *gin.Context) {
	var querySql sql.QuerySql
	err := c.ShouldBindQuery(&querySql)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if requerySql, err := querySqlService.GetQuerySql(querySql.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requerySql": requerySql}, c)
	}
}

// GetQuerySqlList 分页获取QuerySql列表
// @Tags QuerySql
// @Summary 分页获取QuerySql列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query sqlReq.QuerySqlSearch true "分页获取QuerySql列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /querySql/getQuerySqlList [get]
func (querySqlApi *QuerySqlApi) GetQuerySqlList(c *gin.Context) {
	var pageInfo sqlReq.QuerySqlSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := querySqlService.GetQuerySqlInfoList(pageInfo); err != nil {
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
