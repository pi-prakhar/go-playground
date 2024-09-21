package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Single character", "a", "a"},
		{"Two same characters", "aa", "aa"},
		{"Two different characters", "ab", "a"},
		{"Palindrome at start", "abacde", "aba"},
		{"Palindrome at end", "cdeaba", "aba"},
		{"Palindrome in middle", "cbabd", "bab"},
		{"Entire string is palindrome", "racecar", "racecar"},
		{"Multiple palindromes", "cbbd", "bb"},
		{"No palindrome longer than 1", "abcd", "a"},
		{"Empty string", "", ""},
		{"Long string with short palindrome", "abcdefghijklmnopqrstuvwxyz", "a"},
		{"Long palindrome", "abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba", "abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba"},
		{"Odd length palindrome", "abcba", "abcba"},
		{"Even length palindrome", "abccba", "abccba"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("longestPalindrome(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Single character", "a", true},
		{"Two same characters", "aa", true},
		{"Two different characters", "ab", false},
		{"Odd length palindrome", "aba", true},
		{"Even length palindrome", "abba", true},
		{"Not a palindrome", "abcd", false},
		{"Empty string", "", true},
		{"Long palindrome", "abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba", true},
		{"Long non-palindrome", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("isPalindrome(%q) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
