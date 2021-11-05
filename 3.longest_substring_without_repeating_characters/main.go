package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abababacdefc"))
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	testString := []byte(s)

	if len(testString) == 1 {
		return 1
	}

	used := make(map[byte]struct{})

	maxLength := 1

	var maximums []int

next:
	for x := range testString {
		ts := testString[x:]

		for i, v := range ts {
			if i == 0 {
				used[v] = struct{}{}
				continue
			}

			if _, ok := used[v]; !ok {
				maxLength++

				used[v] = struct{}{}
			} else {
				maximums = append(maximums, maxLength)
				maxLength = 1

				used = make(map[byte]struct{})
				continue next
			}
		}
	}

	maximums = append(maximums, maxLength)

	max := 0

	for _, v := range maximums {
		if v > max {
			max = v
		}
	}

	return max
}
