package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	maxIndex := len(cost) - 1
	dp := make([]int, len(cost)+1)
	for i := range dp {
		index := i
		counter := 0
		for {
			if index >= maxIndex {
				break
			}

			if cost[index] >= cost[index+1] && index > i {
				counter += cost[index+1]
				index++
			} else {
				counter += cost[index]
			}

			index++
		}

		dp[i] = counter
	}

	for i := len(cost) - 1; i >= 0; i-- {
		if i == len(cost)-1 {
			dp[i] = cost[i] + dp[i+1]
			continue
		}

		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}

	//fmt.Println(dp)

	return min(dp[0], dp[1])
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
