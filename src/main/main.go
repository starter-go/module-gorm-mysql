package main

import (
	modulegormmysql "github.com/starter-go/module-gorm-mysql"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(nil)
	i.MainModule(modulegormmysql.Module())
	i.WithPanic(true).Run()
}
