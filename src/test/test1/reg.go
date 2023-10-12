package test1

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
)

const (
	theNamespace = "module-gorm-mysql/src/test/test1"
)

// TableReg ...
type TableReg struct {

	//starter:component
	_as func(libgorm.GroupRegistry) //starter:as(".")

	Context application.Context //starter:inject("context")

	agent libgorm.DatabaseAgent
}

func (inst *TableReg) _impl(a libgorm.GroupRegistry) {
	a = inst
}

// Group ...
func (inst *TableReg) Group() *libgorm.Group {

	list := make([]any, 0)
	list = append(list, &Table1{})
	list = append(list, &Table2{})
	list = append(list, &Table3{})

	return &libgorm.Group{
		Enabled:    true,
		Name:       theNamespace,
		Prototypes: list,
		OnInit:     inst.onInit,
		Prefix:     "",
	}
}

func (inst *TableReg) onInit(ctx *libgorm.TableContext) {
	inst.agent.Init(ctx.Database)
}
