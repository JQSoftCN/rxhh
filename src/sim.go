package main

import (
	"../pkg/points"
	"../pkg/sim"
	_"../pkg/tpri"
	"github.com/qjsoftcn/confs"
)

func main() {
	dbName := confs.GetString("top", "RealDB", "Name")
	connUrl := confs.GetString("top", "RealDB", "Url")
	size := confs.GetInt("top", "RealDB", "Size")

	db, err := points.Open(dbName, connUrl, size)
	if err!=nil{
		panic(err)
	}

	sim.MakeSim("kf","dbp",db)
}
