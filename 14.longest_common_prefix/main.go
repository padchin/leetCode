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
	ShortestLength := len(strs[shortestPosition])

	for i, v := range strs {
		if ShortestLength > len(v) {
			ShortestLength = len(v)
			shortestPosition = i
		}
	}

	c := 0

start:
	for i := 0; i < ShortestLength; i++ {
		for j, ss := range strs {
			if j == shortestPosition {
				continue
			}

			if ss[i] != strs[shortestPosition][i] {
				break start
			}
		}

		c++
	}

	return strs[shortestPosition][:c]
}
