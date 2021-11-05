package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

func main() {
	fmt.Println(longestPalindrome(genString(100)))
}

func longestPalindrome(s string) string {
	if utf8.RuneCountInString(s) == 1 {
		return s
	}

	palindrome := ""

	for i := range s {
		ts := s[i:]
	check:
		for x := range ts {
			ss := ts[:x+1] // +1 так как правый символ не включается

			for rev := range ss {
				if ss[rev] != ss[len(ss)-rev-1] {
					continue check
				}
			}

			if len(palindrome) < len(ss) {
				palindrome = ss
			}
		}
	}

	return palindrome
}

func genString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
