package main

import (
	"fmt"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"0A", false},
		{"A man, a plan, a canal: Panama", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Running isPalindrome(%s)", tt.s), func(t *testing.T) {
			got := isPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("isPalindrome(%s) = %t : wanted %t", tt.s, got, tt.want)
			}
		})
	}
}
