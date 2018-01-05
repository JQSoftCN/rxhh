package main

import (
	"../pkg/confs"
	"fmt"
	"time"
)
func main(){

	t1:=time.Now()

	locale,_:=confs.GetString("top","locale")

	val,bool:=confs.GetString("point","Type","Real",locale)

	fmt.Println(val,bool,locale)

	valTypes,bool:=confs.GetStrings("point","ValType")

	for index,v:=range valTypes{
		fmt.Println(index,v)
	}

	fmt.Println(time.Since(t1).Nanoseconds())

}
