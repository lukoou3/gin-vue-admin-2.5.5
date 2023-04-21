package code

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/code"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    codeReq "github.com/flipped-aurora/gin-vue-admin/server/model/code/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ShellcodeApi struct {
}

var shellcodeService = service.ServiceGroupApp.CodeServiceGroup.ShellcodeService


// CreateShellcode 创建Shellcode
// @Tags Shellcode
// @Summary 创建Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body code.Shellcode true "创建Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shellcode/createShellcode [post]
func (shellcodeApi *ShellcodeApi) CreateShellcode(c *gin.Context) {
	var shellcode code.Shellcode
	err := c.ShouldBindJSON(&shellcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    shellcode.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Code":{utils.NotEmpty()},
    }
	if err := utils.Verify(shellcode, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := shellcodeService.CreateShellcode(&shellcode); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShellcode 删除Shellcode
// @Tags Shellcode
// @Summary 删除Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body code.Shellcode true "删除Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shellcode/deleteShellcode [delete]
func (shellcodeApi *ShellcodeApi) DeleteShellcode(c *gin.Context) {
	var shellcode code.Shellcode
	err := c.ShouldBindJSON(&shellcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    shellcode.DeletedBy = utils.GetUserID(c)
	if err := shellcodeService.DeleteShellcode(shellcode); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShellcodeByIds 批量删除Shellcode
// @Tags Shellcode
// @Summary 批量删除Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shellcode/deleteShellcodeByIds [delete]
func (shellcodeApi *ShellcodeApi) DeleteShellcodeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := shellcodeService.DeleteShellcodeByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShellcode 更新Shellcode
// @Tags Shellcode
// @Summary 更新Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body code.Shellcode true "更新Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shellcode/updateShellcode [put]
func (shellcodeApi *ShellcodeApi) UpdateShellcode(c *gin.Context) {
	var shellcode code.Shellcode
	err := c.ShouldBindJSON(&shellcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    shellcode.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Code":{utils.NotEmpty()},
      }
    if err := utils.Verify(shellcode, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := shellcodeService.UpdateShellcode(shellcode); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShellcode 用id查询Shellcode
// @Tags Shellcode
// @Summary 用id查询Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query code.Shellcode true "用id查询Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shellcode/findShellcode [get]
func (shellcodeApi *ShellcodeApi) FindShellcode(c *gin.Context) {
	var shellcode code.Shellcode
	err := c.ShouldBindQuery(&shellcode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshellcode, err := shellcodeService.GetShellcode(shellcode.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshellcode": reshellcode}, c)
	}
}

// GetShellcodeList 分页获取Shellcode列表
// @Tags Shellcode
// @Summary 分页获取Shellcode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query codeReq.ShellcodeSearch true "分页获取Shellcode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shellcode/getShellcodeList [get]
func (shellcodeApi *ShellcodeApi) GetShellcodeList(c *gin.Context) {
	var pageInfo codeReq.ShellcodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shellcodeService.GetShellcodeInfoList(pageInfo); err != nil {
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
