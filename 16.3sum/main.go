package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) //nolint:forbidigo
}

func threeSum(nums []int) [][]int {
	var result [][]int
	usedTriples := make(map[int]struct{})

	for i := range nums {
		twoSumPairs := twoSum(nums, nums[i]*-1, i)

	next:
		for _, pair := range twoSumPairs {
			// формируется тройка
			tripleValues := []int{nums[i]}
			tripleValues = append(tripleValues, pair...)

			sort.Ints(tripleValues)

			p1 := bits.RotateLeft64(uint64(tripleValues[0]), 41)
			p2 := bits.RotateLeft64(uint64(tripleValues[1]), 21)
			p3 := uint64(tripleValues[2])
			s := p1 ^ p2 ^ p3

			//fmt.Printf("%064b\n", s) //nolint:forbidigo

			if _, exists := usedTriples[int(s)]; exists {
				continue next
			}

			result = append(result, tripleValues)
			usedTriples[int(s)] = struct{}{}
		}
	}

	return result
}

// twoSum возвращает массив из пар значений входного массива nums, сумма которых равна target.
func twoSum(nums []int, target int, skip int) [][]int {
	var r [][]int

	used := make(map[int]int)

	for i := range nums {
		if i == skip {
			continue
		}

		c := target - nums[i]

		if _, exists := used[nums[i]]; exists {
			r = append(r, []int{used[nums[i]], nums[i]})
		}

		used[c] = nums[i]
	}

	return r
}
