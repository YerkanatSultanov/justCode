package main

import (
	"fmt"
	"sync"
)

func main() {
	nums1 := make(chan int)
	nums2 := make(chan int)
	nums3 := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			nums1 <- num
		}
		close(nums1)
	}()

	go func() {
		for _, num := range []int{4, 5, 6} {
			nums2 <- num
		}
		close(nums2)
	}()

	go func() {
		for _, num := range []int{7, 8, 9} {
			nums3 <- num
		}
		close(nums3)
	}()

	for num := range merge(nums1, nums2, nums3) {
		fmt.Println(num)
	}
}

func merge(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))

		for _, channel := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for num := range ch {
					mergedCh <- num
				}
			}(channel, wg)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}
