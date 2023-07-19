package driver

import (
	"errors"
	"strconv"
	"strings"

	"github.com/starter-go/libgorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Driver MySQL 驱动
type Driver struct {

	//starter:component
	_as func(libgorm.Driver) //starter:as(".")

}

func (inst *Driver) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *Driver) Registration() *libgorm.DriverRegistration {
	return &libgorm.DriverRegistration{
		Name:   "mysql",
		Driver: inst,
	}
}

// Open ...
func (inst *Driver) Open(cfg *libgorm.Configuration) (libgorm.Database, error) {
	db, err := inst.openDB(cfg)
	if err != nil {
		return nil, err
	}
	builder := &libgorm.DatabaseBuilder{DB: db}
	return builder.Create(), nil
}

func (inst *Driver) openDB(cfg *libgorm.Configuration) (*gorm.DB, error) {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//   dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if cfg == nil {
		return nil, errors.New("config==nil")
	}

	inst.prepareForDefaultPort(cfg)

	dsnbuilder := &strings.Builder{}
	dsnbuilder.WriteString(cfg.User)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(cfg.Password)
	dsnbuilder.WriteString("@tcp(")
	dsnbuilder.WriteString(cfg.Host)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(strconv.Itoa(cfg.Port))
	dsnbuilder.WriteString(")/")
	dsnbuilder.WriteString(cfg.Database)
	dsnbuilder.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := dsnbuilder.String()

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (inst *Driver) prepareForDefaultPort(cfg *libgorm.Configuration) {
	const defport = 3306
	port := cfg.Port
	if port < 1 {
		port = defport
	}
	cfg.Port = port
}
