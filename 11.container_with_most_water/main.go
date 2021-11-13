package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	max := 0

	l := 0
	r := len(height) - 1

	for l < r {
		if height[l] < height[r] {
			v := height[l] * (r - l)

			if max < v {
				max = v
			}

			l++
		} else {
			v := height[r] * (r - l)
			if max < v {
				max = v
			}

			r--
		}
	}

	return max
}
