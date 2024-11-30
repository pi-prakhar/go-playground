package main

import "testing"

func TestGetShortURL(t *testing.T) {
	tests := []struct {
		id     int
		length int
		want   string
	}{
		{id: 0, length: 6, want: "AAAAAA"},
		{id: 1, length: 6, want: "AAAAAB"},
		{id: 56800235583, length: 6, want: "999999"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := getShortURL(6, test.id)
			if test.want != got {
				t.Errorf("wanted %s but got %s", test.want, got)
			}
		})
	}
}
