package main

import (
	"strings"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/a/../../b/../c//.//", "/c"},
		{"/../", "/"},
		{"/../../../", "/"},
		{"//home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := simplifyPath(test.path)
			if strings.Compare(got, test.want) != 0 {
				t.Errorf("want %s but got %s", test.want, got)
			}
		})
	}

}
