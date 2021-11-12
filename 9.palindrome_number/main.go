package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(isPalindrome(2323))
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
		d := int(math.Pow10(t))
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
