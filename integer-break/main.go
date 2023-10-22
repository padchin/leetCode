package main

import "fmt"

var calculated = make(map[int]int)

func p(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	if n == 5 {
		return 6
	}

	if n == 8 {
		return 18
	}

	product := 1
	mod := n % 3
	times := n / 3

	for x := 0; x < times; x++ {
		product = product * 3
	}

	if mod == 1 {
		product = product / 3 * 4
	} else if mod == 2 {
		product = product * 2
	}

	return product
}

func main() {
	fmt.Println(p(10))
}
