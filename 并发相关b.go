package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch = make(chan int)

func main() {
	wg.Add(3)
	go Work("goroutine1")
	<-ch
	go Work("goroutine2")
	<-ch
	go Work("goroutine3")
	<-ch
	wg.Wait()
	fmt.Println("successful")
}
func Work(workName string) {
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	ch <- 1
	wg.Done()
}
