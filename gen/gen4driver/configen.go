package gen4driver

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForGormMySQL ...
func ExportComForGormMySQL(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
