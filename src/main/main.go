package main

import (
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(nil)
	i.MainModule(mysql.Module())
	i.WithPanic(true).Run()
}
