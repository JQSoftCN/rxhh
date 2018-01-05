package main

import (
	"github.com/xormplus/xorm"
	"log"
	_ "github.com/bmizerany/pq"
	"github.com/xormplus/core"
	"fmt"
)

const (
	ConnURL    = "postgres://postgres:224712@localhost/HHSoft?sslmode=disable"
	DriverName = "postgres"
)

var engine xorm.Engine

type Report struct {
	Id          int
	Name        string
	DateType    byte
	DefaultDate string
	DefaultFmt  string
}

func main() {

	engine, err := xorm.NewEngine(DriverName, ConnURL)

	if err != nil {
		log.Println(err)
	}

	err = engine.Ping()

	if err != nil {
		log.Println(err)
	}

	engine.ShowSQL(true)

	engine.SetMaxIdleConns(3)
	engine.SetMaxOpenConns(20)

	tableMapper:=core.NewPrefixMapper(core.SnakeMapper{}, "HH_")
	engine.SetTableMapper(tableMapper)

	engine.CreateTables(&Report{})


	fmt.Println(engine.DBMetas())
	fmt.Println(engine.TableInfo(Report{}))


}
