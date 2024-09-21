package main

import (
	"fmt"
	"testing"
)

func TestGetASCIIValue(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "aaangrm", true},
		{"anagram", "aaamgrk", false},
		{"aaa", "aa", false},
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
		{"hello", "bello", false},
		{"", "", true},    // edge case: both strings are empty
		{"a", "a", true},  // edge case: single character, same
		{"a", "b", false}, // edge case: single character, different
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%s", tt.s, tt.t), func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagram(%s, %s) = %t; want %t", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
