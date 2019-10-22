package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init()  {
	fmt.Println("this is init function. I'm learning about mutex")

}
var wgGlobal sync.WaitGroup
var wgF sync.WaitGroup
var wgB sync.WaitGroup

func main() {

	fmt.Printf("Go Routines:\t %v\n", runtime.NumGoroutine())
	wgGlobal.Add(2)
	wgB.Add(1)
	go foo()
	go bar()
	fmt.Printf("Go Routines:\t %v\n", runtime.NumGoroutine())
	wgGlobal.Wait()

}

func foo() {
	for i := 0; i <= 100; i++ {
		wgF.Wait()
		fmt.Printf("%v \n", i)
		wgB.Done()
		wgF.Add(1)
	}

	wgGlobal.Done()
}

func bar() {
	for i := 100; i >= 0; i-- {
		wgB.Wait()
		fmt.Printf("\t%v \n", i)
		wgF.Done()
		wgB.Add(1)
	}


	wgGlobal.Done()
}
