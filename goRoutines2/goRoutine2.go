package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("this is Go Routine 2")
}

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Go Routines: %v\n", runtime.NumGoroutine())
	//var maxGoRoutines int
	count := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			goCount := count
			goCount++
			count = goCount
			m.Unlock()
			wg.Done()
		}()
		//if runtime.NumGoroutine() > maxGoRoutines {
		//	maxGoRoutines = runtime.NumGoroutine()
		//}
	}
	fmt.Println("Count:",count)
	//fmt.Println("Max Go Routines:",maxGoRoutines)
	wg.Wait()
}
