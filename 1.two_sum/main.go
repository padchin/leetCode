package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 8, 2, -4, 5, 11}, 7))
}

// twoSum возвращает массив из пар значений входного массива nums, сумма которых равна target.
func twoSum(nums []int, target int, skip ...int) [][]int {
	var r [][]int
	used := make(map[int]int)

	for i := range nums {
		if len(skip) > 0 {
			if i == skip[0] {
				continue
			}
		}

		c := target - nums[i]

		if _, exists := used[nums[i]]; exists {
			r = append(r, []int{used[nums[i]], nums[i]})
		}

		used[c] = nums[i]

	}

	return r
}
