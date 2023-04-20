import service from '@/utils/request'

// @Tags Datasource
// @Summary 创建Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Datasource true "创建Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /datasource/createDatasource [post]
export const createDatasource = (data) => {
  return service({
    url: '/datasource/createDatasource',
    method: 'post',
    data
  })
}

// @Tags Datasource
// @Summary 删除Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Datasource true "删除Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /datasource/deleteDatasource [delete]
export const deleteDatasource = (data) => {
  return service({
    url: '/datasource/deleteDatasource',
    method: 'delete',
    data
  })
}

// @Tags Datasource
// @Summary 删除Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /datasource/deleteDatasource [delete]
export const deleteDatasourceByIds = (data) => {
  return service({
    url: '/datasource/deleteDatasourceByIds',
    method: 'delete',
    data
  })
}

// @Tags Datasource
// @Summary 更新Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Datasource true "更新Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /datasource/updateDatasource [put]
export const updateDatasource = (data) => {
  return service({
    url: '/datasource/updateDatasource',
    method: 'put',
    data
  })
}

// @Tags Datasource
// @Summary 用id查询Datasource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Datasource true "用id查询Datasource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /datasource/findDatasource [get]
export const findDatasource = (params) => {
  return service({
    url: '/datasource/findDatasource',
    method: 'get',
    params
  })
}

// @Tags Datasource
// @Summary 分页获取Datasource列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Datasource列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /datasource/getDatasourceList [get]
export const getDatasourceList = (params) => {
  return service({
    url: '/datasource/getDatasourceList',
    method: 'get',
    params
  })
}
