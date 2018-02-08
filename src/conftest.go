package main

import (
	"fmt"
	"time"
	"github.com/qjsoftcn/confs"
)
func main(){

	t1:=time.Now()

	locale:=confs.GetString("top","locale")

	val:=confs.GetString("point","Type","Real",locale)

	fmt.Println(val,locale)

	fmt.Println(time.Since(t1).Nanoseconds())

}
