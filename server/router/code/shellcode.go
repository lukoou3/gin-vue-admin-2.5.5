package code

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShellcodeRouter struct {
}

// InitShellcodeRouter 初始化 Shellcode 路由信息
func (s *ShellcodeRouter) InitShellcodeRouter(Router *gin.RouterGroup) {
	shellcodeRouter := Router.Group("shellcode").Use(middleware.OperationRecord())
	shellcodeRouterWithoutRecord := Router.Group("shellcode")
	var shellcodeApi = v1.ApiGroupApp.CodeApiGroup.ShellcodeApi
	{
		shellcodeRouter.POST("createShellcode", shellcodeApi.CreateShellcode)   // 新建Shellcode
		shellcodeRouter.DELETE("deleteShellcode", shellcodeApi.DeleteShellcode) // 删除Shellcode
		shellcodeRouter.DELETE("deleteShellcodeByIds", shellcodeApi.DeleteShellcodeByIds) // 批量删除Shellcode
		shellcodeRouter.PUT("updateShellcode", shellcodeApi.UpdateShellcode)    // 更新Shellcode
	}
	{
		shellcodeRouterWithoutRecord.GET("findShellcode", shellcodeApi.FindShellcode)        // 根据ID获取Shellcode
		shellcodeRouterWithoutRecord.GET("getShellcodeList", shellcodeApi.GetShellcodeList)  // 获取Shellcode列表
	}
}
