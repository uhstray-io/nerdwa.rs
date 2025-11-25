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

func ReverseListSlow(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	return append([]int{a[len(a)-1]}, ReverseListSlow(a[:len(a)-1])...)
}

func ReverseListInPlace(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
	return a
}

// This is a basic testing function.
// It checks a few cases to ensure the function behaves as expected.
func TestTableDriven_ReverseList(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 4, 4, 4, 5, 6}, []int{6, 5, 4, 4, 4, 2, 1}},
		{[]int{7, 48, 9}, []int{9, 48, 7}},
		{[]int{}, []int{}},
		{[]int{42}, []int{42}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{-1, -2, -3}, []int{-3, -2, -1}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, -1, 0}, []int{0, -1, 1}},
		{[]int{100, 200, 300, 400}, []int{400, 300, 200, 100}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 11, 10}, []int{10, 11, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
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
func Benchmark_ReverseList(b *testing.B) {
	for b.Loop() { // Use b.Loop() for better benchmarking. (Disables compiler optimizations)
		ReverseList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

// Compare all implementations of ReverseList
func BenchmarkAllTableDriven_ReverseList(b *testing.B) {

	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 4, 4, 4, 5, 6}, []int{6, 5, 4, 4, 4, 2, 1}},
		{[]int{7, 48, 9}, []int{9, 48, 7}},
		{[]int{}, []int{}},
		{[]int{42}, []int{42}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{-1, -2, -3}, []int{-3, -2, -1}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, -1, 0}, []int{0, -1, 1}},
		{[]int{100, 200, 300, 400}, []int{400, 300, 200, 100}},
	}

	merges := []struct {
		name string
		fun  func([]int) []int
	}{
		{"ReverseList", ReverseList},
		{"ReverseListSlow", ReverseListSlow},
		{"ReverseListInPlace", ReverseListInPlace},
	}

	for _, merge := range merges {
		for n, test := range tests {
			b.Run(
				fmt.Sprintf("%s/%d", merge.name, n+1),

				func(b *testing.B) {
					for b.Loop() {
						merge.fun(test.input)
					}
				},
			)
		}
	}
}

// Compare all implementations of ReverseList
func BenchmarkAllGenerated_ReverseList(b *testing.B) {

	var tests = []struct {
		input []int
	}{}

	for n := range 8 {
		// Make a table of inputs of size n*n
		input := make([]int, n*n*n)

		// Set input to 0,1,2,3,...,n*n-1
		for i := range input {
			input[i] = i
		}

		// Append to tests
		tests = append(tests, struct {
			input []int
		}{input})
	}

	merges := []struct {
		name string
		fun  func([]int) []int
	}{
		{"ReverseList", ReverseList},
		{"ReverseListSlow", ReverseListSlow},
		{"ReverseListInPlace", ReverseListInPlace},
	}

	for _, merge := range merges {
		for n, test := range tests {
			b.Run(
				fmt.Sprintf("%s/%d", merge.name, n+1),

				func(b *testing.B) {
					for b.Loop() {
						merge.fun(test.input)
					}
				},
			)
		}
	}
}
