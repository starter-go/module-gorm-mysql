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
	_as func(libgorm.TableRegistry) //starter:as(".")

	Context application.Context //starter:inject("context")

}

func (inst *TableReg) _impl() {
	inst._as(inst)
}

// ListTableRegistrations ...
func (inst *TableReg) ListTableRegistrations() []*libgorm.TableRegistration {

	builder := &libgorm.TableGroupBuilder{}
	builder.Group(&libgorm.TableGroup{
		Name:        "test1",
		Namespace:   theNamespace,
		NamerSetter: theNamers.initAll,
	})

	builder.Entity(&Table1{})
	builder.Entity(&Table2{})
	builder.Entity(&Table3{})

	return builder.Create()
}
