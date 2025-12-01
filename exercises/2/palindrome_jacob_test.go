package nerdwars

/// Challange: Implement a Palindrome function and corresponding tests and benchmarks.
///
/// A palindrome is a string that reads the same forwards and backwards.
/// For example, "racecar" and "madam" are palindromes, while "hello" is not.
///
/// The function should take a string as input and return a boolean indicating whether the string is a palindrome.
/// Your

import (
	"fmt"
	"testing"
)

func Palindrome_Jacob(s string) bool {
	// Implement the palindrome check logic here
	return false
}

// This is a basic testing function.
// It checks a few cases to ensure the function behaves as expected.
func TestPalindrome_TableDriven_Jacob(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"Hello", false},
		{"racecar", true},
		{"madam", true},
		{"step on no pets", true},
		{"not a palindrome", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := Palindrome_Jacob(tt.input)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

// This is a basic benchmarking function.
// It measures the time taken to execute the function multiple times.
func BenchmarkPalindrome_Jacob(b *testing.B) {
	for b.Loop() { // Use b.Loop() for better benchmarking. (Disables compiler optimizations)
		Palindrome_Jacob("racecar")
	}
}

// This is a basic testing function.
// It checks a few cases to ensure the function behaves as expected.
func BenchmarkPalindrome_TableDriven_Jacob(b *testing.B) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"Hello", false},
		{"racecar", true},
		{"madam", true},
		{"step on no pets", true},
		{"not a palindrome", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%t", tt.input, tt.expected)

		b.Run(testname, func(b *testing.B) {
			ans := Palindrome_Jacob(tt.input)
			if ans != tt.expected {
				b.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}
