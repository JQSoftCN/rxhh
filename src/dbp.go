package main

import (
	"fmt"
	"../pkg/points"
	_ "../pkg/tpri"
	"time"
)

func main() {

	//fmt.Println(fmt.Sprintf("%.3f", 123.4567))

	tpriDB, err := points.Open("tpri", "admin/admin@16.16.82.195:12084", 5)
	if err != nil {
		fmt.Println(err)
	}

	ps,_:=tpriDB.ReadPoints()


	names := []string{"N7DCS.10MAG01CP301.XG01", "N12FJ1K.H12O1YC2ZHOVC",ps[2].Name,ps[3].Name,ps[4].Name}
	t1 := time.Now()

	pvs, err := tpriDB.ReadSnapshots(names)
	hs := time.Since(t1).Nanoseconds()
	if err != nil {
		fmt.Println(err)
	} else {

		for _, pv := range pvs {
			fmt.Println(pv)
		}

	}

	fmt.Println(hs / 1000 / 1000)

	ps = tpriDB.GetPoints("n5", "", nil, nil)

	for _, p := range ps {
		fmt.Println(p)
	}

	n := ps[4].Name
	s, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-09-12 00:00:00",time.Local)
	e,_ := time.ParseInLocation("2006-01-02 15:04:05", "2017-09-14 10:11:01",time.Local)


	way := points.HistorysComplementWay{}
	way.End = false
	way.Start = false

	//var s8h int64=8*3600

	t2:=time.Now()

	h, err := tpriDB.ReadHistory(n, int32(s.Unix()), int32(e.Unix()), way)

	hs=time.Since(t2).Nanoseconds()

	if err == nil {
		for index, v := range *h.Vals {
			if index==0{
				fmt.Println(index, v)
			}

			if index==len(*h.Vals)-1{
				fmt.Println(index,v)
			}
		}
	} else {

		fmt.Println(err)
	}

	fmt.Println(hs)

	fmt.Println(tpriDB.GetConnUrl(),tpriDB.GetConnSize(),tpriDB.GetDBName())

	t3:=time.Now()
	pvs,err=tpriDB.InterVals(names,int32(e.Unix()),points.IW_Accurate)
	hs= time.Since(t3).Nanoseconds()
	if err!=nil{
		fmt.Println(err)
	}else{
		for index,p:=range pvs{
			fmt.Println(index,p)
		}
	}
	fmt.Println(hs)
}
