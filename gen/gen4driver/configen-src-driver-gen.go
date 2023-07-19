package gen4driver
import (
    p8cf4ceb56 "github.com/starter-go/module-gorm-mysql/driver"
     "github.com/starter-go/application"
)

// type p8cf4ceb56.Driver in package:github.com/starter-go/module-gorm-mysql/driver
//
// id:com-8cf4ceb56f6387fc-driver-Driver
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-Driver
// alias:
// scope:singleton
//
type p8cf4ceb56f_driver_Driver struct {
}

func (inst* p8cf4ceb56f_driver_Driver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8cf4ceb56f6387fc-driver-Driver"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-Driver"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8cf4ceb56f_driver_Driver) new() any {
    return &p8cf4ceb56.Driver{}
}

func (inst* p8cf4ceb56f_driver_Driver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8cf4ceb56.Driver)
	nop(ie, com)

	


    return nil
}


