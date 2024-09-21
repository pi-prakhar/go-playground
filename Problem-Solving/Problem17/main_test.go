package main

import "testing"

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := characterReplacement(test.s, test.k)
			if got != test.want {
				t.Errorf("expected %d but got %d", test.want, got)
			}
		})
	}

}
