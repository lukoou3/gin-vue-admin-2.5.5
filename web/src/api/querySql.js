import service from '@/utils/request'

// @Tags QuerySql
// @Summary 创建QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuerySql true "创建QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /querySql/createQuerySql [post]
export const createQuerySql = (data) => {
  return service({
    url: '/querySql/createQuerySql',
    method: 'post',
    data
  })
}

// @Tags QuerySql
// @Summary 删除QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuerySql true "删除QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /querySql/deleteQuerySql [delete]
export const deleteQuerySql = (data) => {
  return service({
    url: '/querySql/deleteQuerySql',
    method: 'delete',
    data
  })
}

// @Tags QuerySql
// @Summary 删除QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /querySql/deleteQuerySql [delete]
export const deleteQuerySqlByIds = (data) => {
  return service({
    url: '/querySql/deleteQuerySqlByIds',
    method: 'delete',
    data
  })
}

// @Tags QuerySql
// @Summary 更新QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuerySql true "更新QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /querySql/updateQuerySql [put]
export const updateQuerySql = (data) => {
  return service({
    url: '/querySql/updateQuerySql',
    method: 'put',
    data
  })
}

// @Tags QuerySql
// @Summary 用id查询QuerySql
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QuerySql true "用id查询QuerySql"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /querySql/findQuerySql [get]
export const findQuerySql = (params) => {
  return service({
    url: '/querySql/findQuerySql',
    method: 'get',
    params
  })
}

// @Tags QuerySql
// @Summary 分页获取QuerySql列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取QuerySql列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /querySql/getQuerySqlList [get]
export const getQuerySqlList = (params) => {
  return service({
    url: '/querySql/getQuerySqlList',
    method: 'get',
    params
  })
}
