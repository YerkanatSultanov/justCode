package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.RWMutex

func main() {
	go writer()

	for i := 0; i < 10; i++ {
		go reader(i)
	}

	time.Sleep(2 * time.Second)
}

func reader(id int) {
	for {
		mutex.RLock()
		value := count
		mutex.RUnlock()

		fmt.Printf("Reader %d: Value = %d\n", id, value)

		time.Sleep(500 * time.Millisecond)
	}
}

func writer() {
	for {
		mutex.Lock()
		count++
		mutex.Unlock()

		time.Sleep(1 * time.Second)
	}
}
