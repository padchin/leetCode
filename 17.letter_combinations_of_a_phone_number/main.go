package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("79"))
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
	depth := len(d)
	r := make([]string, 0)

	if depth > 4 {
		return r
	}

	if depth > 0 {
		for _, l1 := range p[d[0]] {
			if depth > 1 {
				for _, l2 := range p[d[1]] {
					if depth > 2 {
						for _, l3 := range p[d[2]] {
							if depth > 3 {
								for _, l4 := range p[d[3]] {
									r = append(r, string(l1)+string(l2)+string(l3)+string(l4))
								}
							} else {
								r = append(r, string(l1)+string(l2)+string(l3))
							}
						}
					} else {
						r = append(r, string(l1)+string(l2))
					}
				}
			} else {
				r = append(r, string(l1))
			}
		}
	}

	return r
}
