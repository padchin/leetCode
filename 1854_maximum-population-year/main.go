package main

import "fmt"

func maximumPopulation(logs [][]int) int {
	pd := make(map[int]int)

	minYear := 3000
	maxYear := 0
	maxPersons := 0

	for _, years := range logs {
		if minYear > years[0] {
			minYear = years[0]
		}

		if maxYear < years[1] {
			maxYear = years[1]
		}

		for date := years[0]; date < years[1]; date++ {
			pd[date]++
		}
	}

	y := 0
	for year := minYear; year <= maxYear; year++ {
		count, _ := pd[year]

		if count > maxPersons {
			maxPersons = count
			y = year
		}
	}

	return y
}

// [[2008,2026],[2004,2008],[2034,2035],[1999,2050],[2049,2050],[2011,2035],[1966,2033],[2044,2049]]
func main() {
	fmt.Println(maximumPopulation([][]int{
		{2008, 2026},
		{2004, 2008},
		{2034, 2035},
		{1999, 2050},
		{2049, 2050},
		{2011, 2035},
		{1966, 2033},
		{2044, 2049},
	}))
}
