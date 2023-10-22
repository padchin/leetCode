package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("writing")
		ch <- 1 // block goro;
		fmt.Println("written")
	}()
	fmt.Println("waiting 500 ms")
	time.Sleep(time.Millisecond * 500) // wg
	close(ch)
	for i := range ch {
		fmt.Println(i) // zero-value; @
	}
	time.Sleep(time.Millisecond * 100)
}
