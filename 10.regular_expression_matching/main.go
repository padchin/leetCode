package main

import "fmt"

func toString(mb *[][]bool) string {
	s := ""

	for i := range *mb {
		for j := range (*mb)[i] {
			if (*mb)[i][j] {
				s += "1"
			} else {
				s += "0"
			}
		}

		s += "\n"
	}

	return s
}

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	text := []byte(s)
	pat := []byte(p)

	star := false

	for _, v := range pat {
		if v == '*' {
			star = true

			break
		}
	}

	inLen := len(text)
	patLen := len(pat)

	if !star {
		if inLen != patLen {
			return false
		}
	}

	T := make([][]bool, inLen+1)

	for i := range T {
		T[i] = make([]bool, patLen+1)
	}

	T[0][0] = true

	for i := 1; i <= patLen; i++ {
		if pat[i-1] == '*' {
			T[0][i] = T[0][i-2]
		}
	}

	for i := 1; i <= inLen; i++ {
		for j := 1; j <= patLen; j++ {
			if pat[j-1] == text[i-1] || pat[j-1] == '.' {
				T[i][j] = T[i-1][j-1]

				continue
			}

			if pat[j-1] == '*' {
				T[i][j] = T[i][j-2]
				prev := pat[j-2]

				if prev == text[i-1] || prev == '.' {
					T[i][j] = T[i][j] || T[i-1][j]
				}

				continue
			}
		}
	}

	fmt.Println(toString(&T))

	return T[inLen][patLen]
}
