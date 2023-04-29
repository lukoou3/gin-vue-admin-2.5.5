package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/code"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Sql     sql.RouterGroup
	Code    code.RouterGroup
	Code    code.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
