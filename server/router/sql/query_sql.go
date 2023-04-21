package sql

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuerySqlRouter struct {
}

// InitQuerySqlRouter 初始化 QuerySql 路由信息
func (s *QuerySqlRouter) InitQuerySqlRouter(Router *gin.RouterGroup) {
	querySqlRouter := Router.Group("querySql").Use(middleware.OperationRecord())
	querySqlRouterWithoutRecord := Router.Group("querySql")
	var querySqlApi = v1.ApiGroupApp.SqlApiGroup.QuerySqlApi
	{
		querySqlRouter.POST("createQuerySql", querySqlApi.CreateQuerySql)   // 新建QuerySql
		querySqlRouter.DELETE("deleteQuerySql", querySqlApi.DeleteQuerySql) // 删除QuerySql
		querySqlRouter.DELETE("deleteQuerySqlByIds", querySqlApi.DeleteQuerySqlByIds) // 批量删除QuerySql
		querySqlRouter.PUT("updateQuerySql", querySqlApi.UpdateQuerySql)    // 更新QuerySql
	}
	{
		querySqlRouterWithoutRecord.GET("findQuerySql", querySqlApi.FindQuerySql)        // 根据ID获取QuerySql
		querySqlRouterWithoutRecord.GET("getQuerySqlList", querySqlApi.GetQuerySqlList)  // 获取QuerySql列表
	}
}
