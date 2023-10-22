package main

import (
	"fmt"
	"strings"
)

func convert(in string) string {
	builder := strings.Builder{}

	runes := []rune(in)
	runeCounter := 0
	currentRune := runes[0]

	for _, r := range runes {
		if r == currentRune {
			runeCounter++
			continue
		}

		builder.WriteRune(currentRune)
		builder.WriteString(fmt.Sprint(runeCounter))
		currentRune = r
		runeCounter = 1
	}

	builder.WriteRune(runes[len(runes)-1])
	builder.WriteString(fmt.Sprintf("%d", runeCounter))

	return builder.String()
}

func main() {
	fmt.Println(convert("ASSSDeeeeee@eeeeeeeeeeeeeeeeЧЧЧЧЧЧЧЧЧЧЧЧE"))
}
