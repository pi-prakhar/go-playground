package main

import "testing"

func TestStrStr(t *testing.T) {
	testCases := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Needle found at the beginning",
			haystack: "hello",
			needle:   "he",
			expected: 0,
		},
		{
			name:     "Needle found in the middle",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "Needle found at the end",
			haystack: "hello",
			needle:   "lo",
			expected: 3,
		},
		{
			name:     "Needle not found",
			haystack: "hello",
			needle:   "world",
			expected: -1,
		},
		{
			name:     "Empty needle",
			haystack: "hello",
			needle:   "",
			expected: 0,
		},
		{
			name:     "Needle longer than haystack",
			haystack: "hello",
			needle:   "helloworldwow",
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := strStr(tc.haystack, tc.needle)
			if result != tc.expected {
				t.Errorf("strStr(%q, %q) = %d, expected %d", tc.haystack, tc.needle, result, tc.expected)
			}
		})
	}
}
