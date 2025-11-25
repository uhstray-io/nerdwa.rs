package nerdwars

import (
	"fmt"
	"slices"
	"testing"
)

func ReverseList(a []int) []int {
	b := make([]int, len(a))
	for i := range a {
		b[len(a)-1-i] = a[i]
	}
	return b
}

// This is a basic testing function.
// It checks a few cases to ensure the function behaves as expected.
func TestReverseListTableDriven(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 4, 4, 4, 5, 6}, []int{6, 5, 4, 4, 4, 2, 1}},
		{[]int{7, 8, 9}, []int{9, 8, 7}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			ans := ReverseList(tt.input)
			if !slices.Equal(ans, tt.expected) {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

// This is a basic benchmarking function.
// It measures the time taken to execute the function multiple times.
func BenchmarkReverseList(b *testing.B) {
	for b.Loop() { // Use b.Loop() for better benchmarking. (Disables compiler optimizations)
		ReverseList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
