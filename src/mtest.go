package main

import (
	"time"
	"fmt"
	"github.com/qjsoftcn/texp"
)

func main() {

	tep := texp.NewParser(time.Now(), time.Now())

	base, _ := time.Parse("2006-01-02 15:04:05", "2017-09-02 10:10:10")

	tep.SetBase(base)

	t1 := time.Now()
	t, _ := tep.Parse("o-13m+10d+70mi")

	fmt.Println(t, base.Sub(t1).Seconds())
}
