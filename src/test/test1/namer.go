package test1

import "github.com/starter-go/libgorm"

type namers struct {
	namer     libgorm.TableNamer
	namespace string

	t1 libgorm.TableNameCache
	t2 libgorm.TableNameCache
	t3 libgorm.TableNameCache
}

func (inst *namers) initAll(namer libgorm.TableNamer) {

	inst.namer = namer
	inst.namespace = theNamespace

	inst.initItem(&inst.t1, "t1")
	inst.initItem(&inst.t2, "t2")
	inst.initItem(&inst.t3, "t3")
}

func (inst *namers) initItem(item *libgorm.TableNameCache, simpleName string) {
	item.Init(inst.namer, inst.namespace, simpleName)
}

////////////////////////////////////////////////////////////////////////////////

var theNamers namers

////////////////////////////////////////////////////////////////////////////////

// TableName ...
func (t Table1) TableName() string {
	return theNamers.t1.Name()
}

// TableName ...
func (t Table2) TableName() string {
	return theNamers.t2.Name()
}

// TableName ...
func (t Table3) TableName() string {
	return theNamers.t3.Name()
}
