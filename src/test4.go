package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var counter int = 0

func Count(mu *sync.Mutex) {
	mu.Lock()
	counter++
	fmt.Println("线程执行", counter)
	mu.Unlock()
}

func main() {

	n:=runtime.GOMAXPROCS(2)
	fmt.Println(n,runtime.NumCPU())

	c := make(chan int)

	go wc(c)

	var x = <-c

	fmt.Println(x)

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case d := <-c:
		// 从ch中读取到数据
		fmt.Println("读取到数据了,", d)
	case <-timeout:
		fmt.Println("没有读取到任何数据,超时了")
	}

	go wc(c)
}

func wc(c chan int) {
	c <- 10
}
