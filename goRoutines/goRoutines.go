package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init()  {
	fmt.Println("this is init function.")

}
var wg sync.WaitGroup

func main() {

	fmt.Printf("Go Routines:\t %v\n", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

}

func foo() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%v \n", i)
	}
	wg.Done()
}

func bar() {
	for i := 100; i >= 0; i-- {
		fmt.Printf("\t%v \n", i)
	}
	wg.Done()
}
