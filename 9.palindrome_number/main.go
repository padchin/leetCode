package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(isPalindrome(0))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	n := int(math.Log10(float64(x))) + 1 // находим кол-во цифр в числе

	if n == 1 {
		return true
	}

	buffer := make([]int, n)

	for t := n - 1; t >= 0; t-- {
		d := fastPower(10, t)
		z := x / d

		buffer[n-t-1] = z

		x -= z * d
	}

	s := make([]int, n)
	copy(s, buffer)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return reflect.DeepEqual(s, buffer)
}

// fastPower возвращает x в степени y
func fastPower(x, y int) int {
	switch {
	case y == 0:
		return 1
	case y == 1:
		return x
	case y%2 == 0:
		return fastPower(x*x, y/2)
	default:
		return x * fastPower(x*x, (y-1)/2)
	}
}
