package main

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version1 string
		version2 string
		expected int
	}{
		{"Equal versions", "1.01", "1.001", 0},
		{"First version greater", "1.0.1", "1", 1},
		{"Second version greater", "7.5.2.4", "7.5.3", -1},
		{"Different length, equal", "1.0.0", "1", 0},
		{"Different length, first greater", "1.0.1", "1", 1},
		{"Different length, second greater", "1.0", "1.0.1", -1},
		{"Leading zeros", "1.0", "1.00", 0},
		{"Large version numbers", "999999999.999999999", "999999999.999999998", 1},
		{"Many segments", "1.2.3.4.5.6.7.8", "1.2.3.4.5.6.7.7", 1},
		{"Zero versions", "0.0.0", "0.0.0", 0},
		{"Single digit versions", "1", "2", -1},
		{"Empty string and version", "", "1", -1},
		{"Version and empty string", "1", "", 1},
		{"Both empty strings", "", "", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := compareVersion(tc.version1, tc.version2)
			if result != tc.expected {
				t.Errorf("compareVersion(%q, %q) = %d; want %d", tc.version1, tc.version2, result, tc.expected)
			}
		})
	}
}

func TestGetVersion(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedValue  int
		expectedLength int
	}{
		{"Single digit", "1", 1, 1},
		{"Multiple digits", "123", 123, 3},
		{"With dot", "45.67", 45, 3},
		{"Leading zeros", "001", 1, 3},
		{"Zero", "0", 0, 1},
		{"Empty string", "", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value, length := getVersion(tc.input)
			if value != tc.expectedValue || length != tc.expectedLength {
				t.Errorf("getVersion(%q) = (%d, %d); want (%d, %d)",
					tc.input, value, length, tc.expectedValue, tc.expectedLength)
			}
		})
	}
}
