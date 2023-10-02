package main

import "fmt"

func winnerOfGame(colors string) bool {
	// создаем массив, где каждым элементом будет количество одинаковых смежных рун

	var runes []int

	prevRune := ' '
	count := 1
	firstA := true
	for i, r := range colors {
		if r == prevRune {
			count++
			continue
		}

		prevRune = r
		if i == 0 {
			if r != 'A' {
				// устанавливаем флаг начального цвета (0 - первый B, 1 - первый A)
				firstA = false
			}
			continue
		}
		runes = append(runes, count)
		count = 1

	}
	runes = append(runes, count)

	// вычисляем, сколько ходов может сделать Алиса
	aliceMovesCount := 0
	bobMovesCount := 0
	for i, c := range runes {
		if firstA {
			if i%2 == 0 && c > 2 {
				aliceMovesCount += c - 2
			}
			if i%2 != 0 && c > 2 {
				bobMovesCount += c - 2
			}
		} else {
			if i%2 == 0 && c > 2 {
				bobMovesCount += c - 2
			}
			if i%2 != 0 && c > 2 {
				aliceMovesCount += c - 2
			}
		}
	}

	return aliceMovesCount > bobMovesCount && aliceMovesCount > 0
}

func main() {
	fmt.Println(winnerOfGame("BBBAAAABBBBBBBB"))
}
