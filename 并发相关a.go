package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)
	go func() {
		fmt.Println("出现")
		<-ch
	}()
	ch <- true
}
