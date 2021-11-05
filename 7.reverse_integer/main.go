package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(2147483649))
}

func reverse(x int) int {
	sign := false

	if x < 0 {
		sign = true
		x = -x
	}

	s := strconv.Itoa(x)

	rs := Reverse(s)
	r, err := strconv.Atoi(rs)

	if err != nil {
		return 0
	}

	if r > 2147483647 {
		return 0
	}

	if sign {
		return -r
	}

	return r
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
