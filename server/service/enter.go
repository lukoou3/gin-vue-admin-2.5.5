package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/code"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	SqlServiceGroup     sql.ServiceGroup
	CodeServiceGroup    code.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
