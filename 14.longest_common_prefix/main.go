package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	// find the shortest string
	shortestPosition := 0
	shortestLength := len(strs[shortestPosition])

	for i, v := range strs {
		if shortestLength > len(v) {
			shortestLength = len(v)
			shortestPosition = i
		}
	}

	c := 0

entry:
	for i := 0; i < shortestLength; i++ {
		for j, ss := range strs {
			if j == shortestPosition {
				continue
			}

			if ss[i] != strs[shortestPosition][i] {
				break entry
			}
		}

		c++
	}

	return strs[shortestPosition][:c]
}
