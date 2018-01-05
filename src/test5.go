package main

import (
	"../pkg/reports"
	"../pkg/funcs"
	"../pkg/confs"
	"../pkg/points"
	"fmt"
	"time"
	_"../pkg/tpri"
)

func main() {

	dbName := confs.GetString("top", "RealDB", "Name")
	connUrl := confs.GetString("top", "RealDB", "Url")
	size := confs.GetInt("top", "RealDB", "Size")

	db, err := points.Open(dbName, connUrl, size)
	if err != nil {
		//record err,tell user,use sim replace.
		//this is a good method
		panic(err)
	} else {
		points.SetTopDB(db)
	}

	ParseReport()
	ExcuteReport()
}

func ExcuteReport() {
	rname := "我的报表"

	start := time.Now()
	end := start

	rconfs := reports.GetRConf(rname)
	fctx := funcs.NewFuncContext(start, end)
	//set const
	for index, cc := range *rconfs {
		if cc.IsConst() {
			fctx.AddConst(cc.Key, cc.Cell)
		} else {
			f, ok := cc.Cell.(string)
			if ok {
				fctx.AddFormula(cc.Key, f)
			} else {
				fmt.Println("No.", index, " Formula:", cc, " convert err")
			}
		}
	}

	//calc run
	t1 := time.Now()
	for _, cc := range *rconfs {
		if cc.IsFormula() {
			fctx.RunFormula(cc.Key)
		}
	}

	hs := time.Since(t1).Nanoseconds()

	for _, ce := range *fctx.GetCalcErrs() {
		fmt.Println("err", ce)
	}

	for index, cc := range *rconfs {
		fmt.Println(index, cc.Key, fctx.GetResult(cc.Key))
	}

	fmt.Println("calc time:", hs)
}

func ParseReport() {
	t1 := time.Now()
	path := "C:\\Users\\doudou\\Desktop\\我的报表.xlsx"
	reports.Parse(path)
	hs := time.Since(t1).Nanoseconds()
	fmt.Println(hs)
}
