package gen4test
import (
    p0ef6f2938 "github.com/starter-go/application"
    p0c3ec0b12 "github.com/starter-go/module-gorm-mysql/src/test/test1"
     "github.com/starter-go/application"
)

// type p0c3ec0b12.TableReg in package:github.com/starter-go/module-gorm-mysql/src/test/test1
//
// id:com-0c3ec0b1282d50c6-test1-TableReg
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry
// alias:
// scope:singleton
//
type p0c3ec0b128_test1_TableReg struct {
}

func (inst* p0c3ec0b128_test1_TableReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0c3ec0b1282d50c6-test1-TableReg"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0c3ec0b128_test1_TableReg) new() any {
    return &p0c3ec0b12.TableReg{}
}

func (inst* p0c3ec0b128_test1_TableReg) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0c3ec0b12.TableReg)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)


    return nil
}


func (inst*p0c3ec0b128_test1_TableReg) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


