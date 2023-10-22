package main

import (
	"fmt"
	"sync"
)

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}
}

func joinChannels(chs ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))
		for _, channel := range chs {
			channel := channel

			go func() {
				defer wg.Done()

				for c := range channel {
					merged <- c
				}
			}()
		}

		wg.Wait()
		close(merged)
	}()

	return merged
}
