package main

import (
	"testing"
)

func TestExist(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
		{
			board: [][]byte{
				{'A'},
			},
			word: "A",
			want: true,
		},
		{
			board: [][]byte{
				{'A'},
			},
			word: "B",
			want: false,
		},
		{
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word: "ABCD",
			want: false,
		},
		{
			board: [][]byte{
				{'A', 'D'},
				{'C', 'B'},
			},
			word: "ACBD",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("exist(%v, %s) = %v; want %v", tt.board, tt.word, got, tt.want)
			}
		})
	}
}
