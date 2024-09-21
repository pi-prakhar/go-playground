package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic case",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Multiple spaces",
			input:    "  hello   world  ",
			expected: "world hello",
		},
		{
			name:     "Single word",
			input:    "oneword",
			expected: "oneword",
		},
		{
			name:     "Mixed case",
			input:    "Go Programming Language",
			expected: "Language Programming Go",
		},
		{
			name:     "Special characters",
			input:    "a!b@c#d$e",
			expected: "a!b@c#d$e",
		},
		{
			name:     "Numbers and words",
			input:    "1st 2nd 3rd",
			expected: "3rd 2nd 1st",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverseWords(tc.input)
			if result != tc.expected {
				t.Errorf("reverseWords(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}
