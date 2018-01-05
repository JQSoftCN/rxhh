package main

import (
	"fmt"
	"../pkg/funcs"
	"github.com/bluele/gcache"
	"log"
	"time"
)

func main() {

	str := ":pe('CN7TB.N7TS_T_MSA'+'CN7TB.N7TS_T_MSB'/20,t-1d)/c8+c20"

	str = funcs.Check("s1", str)

	fmt.Println(str)

	gc := gcache.New(100).LRU().LoaderFunc(
		func(key interface{}) (interface{}, error) {

			return str, nil
		}).Build()

	//gc.SetWithExpire("a", str, 5*time.Second)

	v, err := gc.Get("a")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("cache:", v)

	time.Sleep(5 * time.Second)

	v, err = gc.Get("a")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("cache:", v)

}
