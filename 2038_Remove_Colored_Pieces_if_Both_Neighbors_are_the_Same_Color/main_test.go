package main

import (
	"fmt"
	"testing"
)

func TestWinnerOfGame(t *testing.T) {
	testCases := []struct {
		board    string
		expected bool
	}{
		{"AAABABB", true},
		{"AA", false},
		{"ABBBBBBBAAA", false},
		{"AAAABBBB", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Board: %s", tc.board), func(t *testing.T) {
			result := winnerOfGame(tc.board)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
