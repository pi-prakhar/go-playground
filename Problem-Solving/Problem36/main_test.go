package main

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{"n = 1", 1, "1"},
		{"n = 2", 2, "11"},
		{"n = 3", 3, "21"},
		{"n = 4", 4, "1211"},
		{"n = 5", 5, "111221"},
		{"n = 6", 6, "312211"},
		{"n = 7", 7, "13112221"},
		{"n = 8", 8, "1113213211"},
		{"n = 9", 9, "31131211131221"},
		{"n = 10", 10, "13211311123113112211"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countAndSay(tc.input)
			if result != tc.expected {
				t.Errorf("countAndSay(%d) = %s; want %s", tc.input, result, tc.expected)
			}
		})
	}
}

func TestGetRLEOfNumber(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Single digit", "1", "11"},
		{"Two same digits", "11", "21"},
		{"Two different digits", "12", "1112"},
		{"Multiple same digits", "111", "31"},
		{"Multiple different digits", "1234", "11121314"},
		{"Mixed repeated digits", "11223344", "21222324"},
		{"Complex case", "3322251", "23321511"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getRLEOfNumber(tc.input)
			if result != tc.expected {
				t.Errorf("getRLEOfNumber(%s) = %s; want %s", tc.input, result, tc.expected)
			}
		})
	}
}

func BenchmarkCountAndSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countAndSay(30)
	}
}

func BenchmarkCountAndSayOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countAndSayOptimized(30)
	}
}
