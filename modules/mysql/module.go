package mysql

import (
	"github.com/starter-go/application"
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	"github.com/starter-go/module-gorm-mysql/gen/gen4driver"
)

// Module 函数返回 mysql 的 libgorm 驱动模块
func Module() application.Module {
	mb := modulegormmysql.NewDriverModule()
	mb.Components(gen4driver.ExportComForGormMySQL)
	return mb.Create()
}
