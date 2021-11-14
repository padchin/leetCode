package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("234"))
}

func letterCombinations(digits string) []string {
	p := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	d := []byte(digits)

	for _, v := range d {
		for _, letters := range p[v] {

		}
	}
}
