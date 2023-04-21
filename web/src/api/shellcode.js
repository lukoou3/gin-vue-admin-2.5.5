import service from '@/utils/request'

// @Tags Shellcode
// @Summary 创建Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shellcode true "创建Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shellcode/createShellcode [post]
export const createShellcode = (data) => {
  return service({
    url: '/shellcode/createShellcode',
    method: 'post',
    data
  })
}

// @Tags Shellcode
// @Summary 删除Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shellcode true "删除Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shellcode/deleteShellcode [delete]
export const deleteShellcode = (data) => {
  return service({
    url: '/shellcode/deleteShellcode',
    method: 'delete',
    data
  })
}

// @Tags Shellcode
// @Summary 删除Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shellcode/deleteShellcode [delete]
export const deleteShellcodeByIds = (data) => {
  return service({
    url: '/shellcode/deleteShellcodeByIds',
    method: 'delete',
    data
  })
}

// @Tags Shellcode
// @Summary 更新Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Shellcode true "更新Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shellcode/updateShellcode [put]
export const updateShellcode = (data) => {
  return service({
    url: '/shellcode/updateShellcode',
    method: 'put',
    data
  })
}

// @Tags Shellcode
// @Summary 用id查询Shellcode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Shellcode true "用id查询Shellcode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shellcode/findShellcode [get]
export const findShellcode = (params) => {
  return service({
    url: '/shellcode/findShellcode',
    method: 'get',
    params
  })
}

// @Tags Shellcode
// @Summary 分页获取Shellcode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Shellcode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shellcode/getShellcodeList [get]
export const getShellcodeList = (params) => {
  return service({
    url: '/shellcode/getShellcodeList',
    method: 'get',
    params
  })
}
