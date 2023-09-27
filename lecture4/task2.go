package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	resultCh := make(chan int, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()
			value, ok := m["key"]
			if ok {
				m["key"] = value + 1
			} else {
				m["key"] = 1
			}
			resultCh <- m["key"]
			mu.Unlock()
		}()
	}

	wg.Wait()
	close(resultCh)

	for result := range resultCh {
		fmt.Println(result)
	}

}
