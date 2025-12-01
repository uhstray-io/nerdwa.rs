package nerdwars

import (
	"fmt"
	"testing"
)

func IntMin_Alize(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// This is a basic testing function.
// It checks a few cases to ensure the function behaves as expected.
func TestIntMinTableDriven_Alize(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin_Alize(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// This is a basic benchmarking function for IntMin.
// It measures the time taken to execute the function multiple times.
func BenchmarkIntMin_Alize(b *testing.B) {
	for b.Loop() { // Use b.Loop() for better benchmarking. (Disables compiler optimizations)
		IntMin_Alize(1, 2)
	}
}
