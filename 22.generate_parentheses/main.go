package main

import (
	"errors"
	"fmt"
	"math/bits"
)

var IllegalPatternErr = errors.New("illegal pattern")

func main() {
	fmt.Println(generateParenthesis(1))
}

func generateParenthesis(size int) []string {
	if size == 0 {
		return []string{""}
	}

	var patterns []string

	// вычисляем первое число 11110000
	first := 0

	for i := size*2 - 1; i > size-1; i-- {
		first |= 1 << i
	}

	// вычисляем последнее число 10101010
	last := 0

	for i := 1; i < (size * 2); i += 2 {
		last |= 1 << i
	}

	for i := first; i >= last; i -= 2 {
		s, err := stringFromBits(i, size)

		if err != nil {
			continue
		}

		patterns = append(patterns, s)
	}

	return patterns
}

func stringFromBits(num int, size int) (string, error) {
	const LEFT = '('
	const RIGHT = ')'

	if bits.OnesCount(uint(num)) != size {
		return "", IllegalPatternErr
	}

	bitsWidth := size * 2

	buffer := make([]byte, bitsWidth)

	for i := bitsWidth - 1; i >= 0; i-- {
		a := bitsWidth - 1 - i

		if num&(1<<i) != 0 {
			buffer[a] = LEFT
		} else {
			buffer[a] = RIGHT
		}
	}

	leftCount := 0
	rightCount := 0

	for _, s := range buffer {
		if s == LEFT {
			leftCount++
		}

		if s == RIGHT {
			rightCount++
		}

		if rightCount > leftCount {
			return "", IllegalPatternErr
		}
	}

	return string(buffer), nil
}
