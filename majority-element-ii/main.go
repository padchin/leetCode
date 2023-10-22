package main

func majorityElement(nums []int) []int {
	var res []int
	numbers := make(map[int]int)

	repeatThreshold := len(nums) / 3

	for _, n := range nums {
		numbers[n]++
	}

	for k, v := range numbers {
		if v > repeatThreshold {
			res = append(res, k)
		}
	}

	return res
}
