package main

import "fmt"

func add(first []int, second []int) []int {
	if len(first) > len(second) {
		zeroes := make([]int, len(first)-len(second))
		second = append(zeroes, second...)
	} else if len(first) < len(second) {
		zeroes := make([]int, len(second)-len(first))
		first = append(zeroes, first...)
	}

	res := make([]int, len(first)+1)

	for i := len(first) - 1; i >= 0; i-- {
		sum := first[i] + second[i]

		if sum > 9 {
			sum -= 10
			if i > 0 {
				first[i-1]++
			} else {
				res[i]++
			}
		}
		res[i+1] = sum
	}

	return res
}

func main() {
	fmt.Println(add([]int{9, 9, 9, 9, 9, 9, 8, 8, 7}, []int{2, 2, 2, 2, 2, 0, 0}))
}
