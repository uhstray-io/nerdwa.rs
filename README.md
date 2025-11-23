# nerdwars

This repository is a benchmarking and testing framework for comparing multiple implementations of a function in Go. The goal is to provide a simple, extensible setup for writing, testing, and benchmarking different solutions to the same problem.

## How it works

- Implement your function(s) in Go.
- Add table-driven tests to verify correctness.
- Add benchmarks to measure performance.
- Use the provided commands to run tests and benchmarks, and to profile CPU and memory usage.

## Running Tests 

Run all tests:

```bash
go test
```
Run tests with verbose output:

```bash
go test -v
```

Run a specific test function:

```bash
go test -run=TestIntMinTableDriven
```

## Benchmarks

Run all benchmarks:

```bash
go test -bench=.
```

Run a specific benchmark:

```bash
go test -bench=BenchmarkIntMin
```

## Profiling

Run a benchmark with CPU profiling:
You will need graphviz installed to visualize the profiles.
https://graphviz.gitlab.io/download/

```bash
go test -bench=BenchmarkIntMin -cpuprofile=cpu.prof
```

Run a benchmark with memory profiling:

```bash
go test -bench=BenchmarkIntMin -memprofile=mem.prof
```

View CPU profile:

```bash
go tool pprof cpu.prof
```

View memory profile:

```bash
go tool pprof mem.prof
```

Run the benchmark with memory allocation statistics (without running tests):

```bash
go test -benchmem -run=^$ -bench ^BenchmarkIntMin$ nerdwa.rs
```


## Contributing

1. Add your function implementation and corresponding tests/benchmarks.
2. Follow the table-driven test and benchmark patterns in `intmin_test.go`.
3. Submit a pull request with your changes.

## License

MIT