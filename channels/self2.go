package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	ints1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ints2 := []int{10, 20, 30, 40}
	ints3 := []int{100, 200, 300, 400, 500, 600}

	go func() {
		for _, v := range ints1 {
			a <- v
		}

		close(a)
	}()

	go func() {
		for _, v := range ints2 {
			b <- v
		}

		close(b)
	}()

	go func() {
		for _, v := range ints3 {
			c <- v
		}

		close(c)
	}()

	for d := range join1(a, b, c) {
		fmt.Println(d)
	}
}

func join1(channels ...chan int) chan int {
	merged := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(channels))

		for _, channel := range channels {
			go func(ch chan int) {
				defer wg.Done()

				for c := range ch {
					merged <- c
				}
			}(channel)
		}

		wg.Wait()
		close(merged)
	}()

	return merged
}
