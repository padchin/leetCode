package main

import (
	"errors"
	"math/bits"
)

var IllegalPatternErr = errors.New("illegal pattern")

func generateParenthesisBits(size int) []string {
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

// Функция generateParenthesis создает и возвращает массив со всеми возможными правильными скобочными последовательностями длины 2*n
func generateParenthesisRecursion(size int) []string {
	// Создаем пустой массив строк для хранения правильных скобочных последовательностей
	var ans []string
	// Вызываем функцию backtrack, чтобы заполнить массив ans
	backtrack(&ans, "", 0, 0, size)
	// Возвращаем массив правильных скобочных последовательностей
	return ans
}

// Функция backtrack генерирует правильные скобочные последовательности рекурсивно
func backtrack(ans *[]string, cur string, open, close, max int) {
	// Если длина текущей последовательности равна 2*n, то добавляем ее в массив ans
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	// Если количество открывающих скобок меньше n, то добавляем открывающую скобку и вызываем backtrack рекурсивно
	if open < max {
		backtrack(ans, cur+"(", open+1, close, max)
	}
	// Если количество закрывающих скобок меньше количества открывающих скобок, то добавляем закрывающую скобку и вызываем backtrack рекурсивно
	if close < open {
		backtrack(ans, cur+")", open, close+1, max)
	}
}
