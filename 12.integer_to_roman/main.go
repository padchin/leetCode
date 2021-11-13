package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(intToRoman(1979))
}

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}

	s := ""

	L := map[int][2]string{
		1: {"I", "V"},
		2: {"X", "L"},
		3: {"C", "D"},
		4: {"M", ""},
	}

	// parse int
	n := int(math.Log10(float64(num))) + 1 // находим кол-во цифр в числе

	buffer := make([]int, n)

	for t := n - 1; t >= 0; t-- {
		d := int(math.Pow10(t))
		z := num / d

		buffer[t] = z

		num -= z * d
	}

	for i := n; i > 0; i-- {
		d := buffer[i-1]

		switch {
		case d == 9:
			s += L[i][0] + L[i+1][0]
		case d == 4:
			s += L[i][0] + L[i][1]
		case d >= 5:
			s += L[i][1]

			for j := 0; j < d-5; j++ {
				s += L[i][0]
			}
		default:
			for j := 0; j < d; j++ {
				s += L[i][0]
			}
		}
	}

	return s
}
