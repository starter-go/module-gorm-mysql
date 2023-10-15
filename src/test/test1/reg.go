package test1

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

const (
	theGroupURI = "uri:module-gorm-mysql/src/test/test1"
)

// TableReg ...
type TableReg struct {

	//starter:component
	_as func(libgorm.GroupRegistry) //starter:as(".")

	Context application.Context       //starter:inject("context")
	DSMan   libgorm.DataSourceManager //starter:inject("#")

	agent libgorm.DataSourceAgent
}

func (inst *TableReg) _impl(a libgorm.GroupRegistry, b libgorm.Group, c libgorm.Agent) {
	a = inst
	b = inst
	c = inst
}

// Groups ...
func (inst *TableReg) Groups() []*libgorm.GroupRegistration {
	r1 := &libgorm.GroupRegistration{
		Enabled: true,
		Alias:   "",
		URI:     theGroupURI,
		Prefix:  "",
		Source:  "",

		Group: inst,
	}
	return []*libgorm.GroupRegistration{r1}
}

// Prototypes ...
func (inst *TableReg) Prototypes() []any {
	list := make([]any, 0)
	list = append(list, &Table1{})
	list = append(list, &Table2{})
	list = append(list, &Table3{})
	return list
}

// DB ...
func (inst *TableReg) DB(db *gorm.DB) *gorm.DB {
	if !inst.agent.Ready() {
		inst.agent.Init(inst.DSMan, "main")
	}
	return inst.agent.DB(db)
}
