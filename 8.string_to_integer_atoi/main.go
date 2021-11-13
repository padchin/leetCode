package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("20000000000000000000"))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	inBytes := []byte(s)

	var filtered []byte

	// убираем первые пробелы
	for i, v := range inBytes {
		if v == ' ' {
			continue
		}

		filtered = inBytes[i:]
		break
	}

	plus := false
	minus := false

	if len(filtered) == 0 {
		return 0
	}

	// поиск знака, установка флага знака и удаление лишнего символа знака
	if filtered[0] == '+' {
		plus = true
	}

	if filtered[0] == '-' {
		minus = true
	}

	if plus || minus {
		filtered = filtered[1:]
	}

	// читаем оставшуюся строку до первого нецифрового символа или до конца ввода
	endPosition := 0
	found := false

	for i, v := range filtered {
		endPosition = i

		if v < '0' || v > '9' {
			found = true

			break
		}
	}

	if found {
		filtered = filtered[:endPosition]
	}

	if len(filtered) == 0 {
		return 0
	}

	i, err := strconv.Atoi(string(filtered))

	if err != nil {
		if minus { // FIXME костыль
			return -2147483648
		}

		return 2147483647
	}

	if minus {
		i = -i
	}

	if i < -2147483648 {
		i = -2147483648
	}

	if i > 2147483647 {
		i = 2147483647
	}

	return i
}
