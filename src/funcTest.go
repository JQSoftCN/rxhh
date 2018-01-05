package main

import (
	"log"
	"../pkg/funcs"
	"fmt"
)

func main() {

	funcCtx := funcs.DefaultFuncContext()

	cells := make(map[string]string)

	cells["A1"] = "100"
	cells["A2"] = "200"
	cells["A3"] = "pe(\"'tagA'\",'t')"
	cells["S1.A4"] = "pe(\"'tagA'+'tagb'+'tagc'\",'t-1d-1s')+2*pe(\"'tagA'+'tagb'+'tagc'\")-300"
	cells["A8"]="'开始'"

	//$1

	for k, v := range cells {

		ret, err := funcCtx.GetVM().Run(v)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(k, v, ret.String())
	}

	fmt.Println()

	for k, v := range *(funcCtx.GetResults().GetMap()) {
		fmt.Println("key:", k, "val:", v)
	}

}
