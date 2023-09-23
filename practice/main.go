package main

import "fmt"

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		for num := range nums {
			squares <- num * num
		}
		close(squares)
	}()

	go func() {
		for _, num := range []int{1, 2, 3, 4, 5} {
			nums <- num
		}
		close(nums)
	}()

	for num := range squares {
		fmt.Println(num)
	}
}
