package main

import "testing"

func Test_exists(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			board: [][]byte{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}},
			word:  "ABCD",
			want:  true,
		}, {
			board: [][]byte{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}},
			word:  "SEE",
			want:  true,
		}, {
			board: [][]byte{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}},
			word:  "ABDE",
			want:  true,
		},
	}
}
