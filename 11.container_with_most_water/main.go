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
			if max < height[l]*(r-l) {
				max = height[l] * (r - l)
			}

			l++
		} else {
			if max < height[r]*(r-l) {
				max = height[r] * (r - l)
			}

			r--
		}
	}

	return max
}
