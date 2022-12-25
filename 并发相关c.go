package main

import (
	"fmt"
	"sync"
)

var csh = make(chan int)
var wg2 sync.WaitGroup

func main() {
	wg2.Add(6)
	go sss1()
	<-csh
	go sss2()
	<-csh
	go sss3()
	<-csh
	go sss4()
	<-csh
	go sss5()
	<-csh
	go sss6()
	<-csh
	wg2.Wait()
	fmt.Println("successful")
}
func sss1() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	csh <- 1
	wg2.Done()
}
func sss2() {
	for i := 10; i < 100; i++ {
		a := i / 10
		b := i % 10
		if i == a*a+b*b {
			fmt.Println(i)
		}
	}
	csh <- 1
	wg2.Done()
}
func sss3() {
	for i := 100; i < 1000; i++ {
		a := i / 100
		b := i % 10
		c := (i - a*100 - b) / 10
		if i == a*a*a+b*b*b+c*c*c {
			fmt.Println(i)
		}
	}
	csh <- 1
	wg2.Done()
}
func sss4() {
	for i := 1000; i < 10000; i++ {
		d := i / 1000
		a := (i - d*1000) / 100
		b := i % 10
		c := (i - d*1000 - a*100) / 10

		if i == a*a*a*a+b*b*b*b+c*c*c*c+d*d*d*d {
			fmt.Println(i)
		}
	}
	csh <- 1
	wg2.Done()
}
func sss5() {
	for i := 10000; i < 100000; i++ {
		e := i / 10000
		d := (i - e*10000) / 1000
		a := (i - e*10000 - d*1000) / 100
		b := i % 10
		c := (i - e*10000 - d*1000 - a*100) / 10

		if i == a*a*a*a*a+b*b*b*b*b+c*c*c*c*c+d*d*d*d*d+e*e*e*e*e {
			fmt.Println(i)
		}
	}
	csh <- 1
	wg2.Done()
}
func sss6() {
	for i := 100000; i < 1000000; i++ {
		f := i / 100000
		e := (i - f*100000) / 10000
		d := (i - f*100000 - e*10000) / 1000
		a := (i - f*100000 - e*10000 - d*1000) / 100
		b := i % 10
		c := (i - f*100000 - e*10000 - d*1000 - a*100) / 10

		if i == a*a*a*a*a*a+b*b*b*b*b*b+c*c*c*c*c*c+d*d*d*d*d*d+e*e*e*e*e*e+f*f*f*f*f*f {
			fmt.Println(i)
		}
	}
	csh <- 1
	wg2.Done()
}
