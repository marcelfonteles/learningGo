package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("Learning about race condition")
}

func main() {
	fmt.Printf("CPUs:\t\t %v\n", runtime.NumCPU())
	fmt.Printf("Go Routines:\t %v\n", runtime.NumGoroutine())

	counter := 0
	var maxGoRoutines int

	const GS = 100
	var wg sync.WaitGroup
	wg.Add(GS)
	var mut sync.Mutex

	for i := 0; i < GS; i++ {
		go func() {
			mut.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mut.Unlock()
			wg.Done()
		}()
		if runtime.NumGoroutine() > maxGoRoutines {
			maxGoRoutines = runtime.NumGoroutine()
		}
	}
	wg.Wait()
	fmt.Printf("Go Routines:\t %v\n", runtime.NumGoroutine())
	fmt.Println("count:",counter)
	fmt.Println("Max Go Routines:",maxGoRoutines)
}
