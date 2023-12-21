package modulegormmysql

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
)

const (
	theModuleName     = "github.com/starter-go/module-gorm-mysql"
	theModuleVersion  = "v1.0.0"
	theModuleRevision = 6
	theModuleResPath  = "src/lib/resources"
)

//go:embed "src/lib/resources"
var theModuleResFS embed.FS

// NewDriverModule 导出模块 ['github.com/starter-go/module-gorm-mysql']
func NewDriverModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Depend(libgorm.Module())

	return mb
}
