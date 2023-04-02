package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(fourSumExperimental([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {

}

func fourSumBruteForce(nums []int, target int) [][]int {
	result := make([][]int, 0)
	used := make(map[int]struct{})

	for i1 := 0; i1 < len(nums); i1++ {
		for i2 := i1 + 1; i2 < len(nums); i2++ {
			for i3 := i2 + 1; i3 < len(nums); i3++ {
				for i4 := i3 + 1; i4 < len(nums); i4++ {
					if nums[i1]+nums[i2]+nums[i3]+nums[i4] == target {
						fourValues := []int{nums[i1], nums[i2], nums[i3], nums[i4]}
						sort.Ints(fourValues)

						p1 := bits.RotateLeft64(uint64(fourValues[0]), 48)
						p2 := bits.RotateLeft64(uint64(fourValues[1]), 32)
						p3 := bits.RotateLeft64(uint64(fourValues[2]), 16)
						p4 := uint64(fourValues[3])
						s := p1 ^ p2 ^ p3 ^ p4

						if _, exists := used[int(s)]; exists {
							continue
						}

						result = append(result, fourValues)
						used[int(s)] = struct{}{}
					}
				}
			}
		}
	}

	return result
}

func fourSumExperimental(nums []int, target int) [][]int {
	result := make([][]int, 0)
	used := make(map[int]struct{})

	for i := range nums {
		threeSumTriples := threeSum(nums, i)

	next:
		for _, triple := range threeSumTriples {
			fourValues := make([]int, 0)

			if nums[i] == -target {
				fourValues = []int{nums[i]}
				fourValues = append(fourValues, triple...)
			}

			if len(fourValues) > 0 {
				sort.Ints(fourValues)

				s := hashArrayInts64(&fourValues)

				if _, exists := used[s]; exists {
					continue next
				}

				result = append(result, fourValues)
				used[s] = struct{}{}
			}
		}
	}

	return result
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

func threeSum(nums []int, skip ...int) [][]int {
	result := make([][]int, 0)
	usedTriples := make(map[int]struct{})

	for i := range nums {
		if len(skip) > 0 {
			if i == skip[0] {
				continue
			}
		}

		twoSumPairs := twoSum(nums, nums[i]*-1, i)

	next:
		for _, pair := range twoSumPairs {
			// формируется тройка
			tripleValues := []int{nums[i]}
			tripleValues = append(tripleValues, pair...)

			sort.Ints(tripleValues)

			s := hashArrayInts64(&tripleValues)

			if _, exists := usedTriples[s]; exists {
				continue next
			}

			result = append(result, tripleValues)
			usedTriples[s] = struct{}{}
		}
	}

	return result
}

// hashArrayInts64 returns hash for first 64 elements of array of ints
func hashArrayInts64(arr *[]int) int {
	parts := len(*arr)
	shift := 64 / parts

	if len(*arr) > 64 {
		shift = 1
	}

	res := uint64(0)

	for i, v := range *arr {
		if i == 64 {
			return int(res)
		}

		p := bits.RotateLeft64(uint64(v), shift*i)
		res ^= p
	}

	return int(res)
}
