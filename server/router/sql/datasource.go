package sql

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DatasourceRouter struct {
}

// InitDatasourceRouter 初始化 Datasource 路由信息
func (s *DatasourceRouter) InitDatasourceRouter(Router *gin.RouterGroup) {
	datasourceRouter := Router.Group("datasource").Use(middleware.OperationRecord())
	datasourceRouterWithoutRecord := Router.Group("datasource")
	var datasourceApi = v1.ApiGroupApp.SqlApiGroup.DatasourceApi
	{
		datasourceRouter.POST("createDatasource", datasourceApi.CreateDatasource)   // 新建Datasource
		datasourceRouter.DELETE("deleteDatasource", datasourceApi.DeleteDatasource) // 删除Datasource
		datasourceRouter.DELETE("deleteDatasourceByIds", datasourceApi.DeleteDatasourceByIds) // 批量删除Datasource
		datasourceRouter.PUT("updateDatasource", datasourceApi.UpdateDatasource)    // 更新Datasource
	}
	{
		datasourceRouterWithoutRecord.GET("findDatasource", datasourceApi.FindDatasource)        // 根据ID获取Datasource
		datasourceRouterWithoutRecord.GET("getDatasourceList", datasourceApi.GetDatasourceList)  // 获取Datasource列表
	}
}
