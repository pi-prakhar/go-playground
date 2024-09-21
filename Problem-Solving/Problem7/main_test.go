package main

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"au", 2},
		{" ", 1},
		{"", 0},
		{"aab", 2},
		{"bpfbhmipx", 7},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("lengthOfLongestSubstring(%s)", tt.s), func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%s) = (%d) want (%d)", tt.s, got, tt.want)
			}
		})
	}
}
