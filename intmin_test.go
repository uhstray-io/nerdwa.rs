package nerdwars

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// This is a basic testing function for IntMin.
// It checks a few cases to ensure the function behaves as expected.

func TestIntMinTableDriven(t *testing.T) {
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
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// This is a basic benchmarking function for IntMin.
// It measures the time taken to execute the function multiple times.
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() { // Use b.Loop() for better benchmarking. (Disables compiler optimizations)
		IntMin(1, 2)
	}
}

// To run the benchmark, use the command:
// go test -bench=.

// To run just the tests, use the command:
// go test

// To run just the selected benchmark, use the command:
// go test -bench=BenchmarkIntMin

// To run the benchmark with CPU profiling, use the command:
// go test -bench=BenchmarkIntMin -cpuprofile=cpu.prof

// To run the benchmark with memory profiling, use the command:
// go test -bench=BenchmarkIntMin -memprofile=mem.prof

// To view the CPU profile, use the command:
// go tool pprof cpu.prof

// To view the memory profile, use the command:
// go tool pprof mem.prof

// The -benchmem flag can be added to any of the above commands to include memory allocation statistics in the benchmark results.

// The below command will run the benchmark with memory allocation statistics but will not run any tests:
// The outrput will show the number of bytes allocated and the number of allocations per operation.
// go test -benchmem -run=^$ -bench ^BenchmarkIntMin$ nerdwa.rs

// To run the tests with verbose output, use the command:
// go test -v
