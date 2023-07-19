package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	"github.com/starter-go/module-gorm-mysql/gen/gen4test"
	"github.com/starter-go/starter"
)

func main() {

	m := module()

	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed "resources"
var theResFS embed.FS

func module() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("module-gorm-mysql/src/test")
	mb.Version("v1")
	mb.Revision(1)

	mb.EmbedResources(theResFS, "resources")
	mb.Components(gen4test.ExportComForGormMySqlTest)
	mb.Depend(modulegormmysql.Module())

	return mb.Create()
}
