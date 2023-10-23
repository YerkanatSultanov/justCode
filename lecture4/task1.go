package main

import (
	"fmt"
	"sync"
)

//race condition example with just counter

func main() {
	var wg sync.WaitGroup
	var count int

	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final count: ", count)
}
