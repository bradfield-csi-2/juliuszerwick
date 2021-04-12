# Memory Hiearchy 2 Prework - Optimizing a simple metrics program

## Initial run of tests with benchmarks prior to optimizations

<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               180           6965447 ns/op
BenchmarkMetrics/Average_payment-4            30          41430489 ns/op
BenchmarkMetrics/Payment_stddev-4             15          79675550 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  9.374s


## 
