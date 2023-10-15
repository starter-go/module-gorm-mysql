package modulegormmysql

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/gen/gen4driver"
)

const (
	theModuleName     = "github.com/starter-go/module-gorm-mysql"
	theModuleVersion  = "v0.9.2"
	theModuleRevision = 3
	theModuleResPath  = "src/lib/resources"
)

//go:embed "src/lib/resources"
var theModuleResFS embed.FS

// Module 导出模块 ['github.com/starter-go/module-gorm-mysql']
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)

	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4driver.ExportComForGormMySQL)

	mb.Depend(libgorm.Module())

	return mb.Create()
}
