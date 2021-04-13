# Memory Hiearchy 2 Prework - Optimizing a simple metrics program

## Initial run of tests with benchmarks prior to optimizations
```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               180           6965447 ns/op
BenchmarkMetrics/Average_payment-4            30          41430489 ns/op
BenchmarkMetrics/Payment_stddev-4             15          79675550 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  9.374s
```

## Changing data types of fields in structs didn't have a significant impact.
```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               178           6772249 ns/op
BenchmarkMetrics/Average_payment-4            33          35431727 ns/op
BenchmarkMetrics/Payment_stddev-4             16          69156423 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  9.475s
```

## Changing function body of AverageAge to iterate over a slice of ages
```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               169           7627713 ns/op
BenchmarkMetrics/Average_payment-4            34          35836195 ns/op
BenchmarkMetrics/Payment_stddev-4             15          70104574 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  9.058s
```

```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               166           7346414 ns/op
BenchmarkMetrics/Average_payment-4            31          35641680 ns/op
BenchmarkMetrics/Payment_stddev-4             16          68105932 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  10.170s
```

```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               164           7485426 ns/op
BenchmarkMetrics/Average_payment-4            34          36029133 ns/op
BenchmarkMetrics/Payment_stddev-4             18          69756435 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  9.636s
```


## Changing function body of AverageAge to iterate over an array of ages

```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               180           6707943 ns/op
BenchmarkMetrics/Average_payment-4            34          35080970 ns/op
BenchmarkMetrics/Payment_stddev-4             18          67713490 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  9.450s
```

```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               177           7039619 ns/op
BenchmarkMetrics/Average_payment-4            30          35405423 ns/op
BenchmarkMetrics/Payment_stddev-4             18          68830893 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  8.568s
```

```
<metrics>$ go test -bench=.
goos: darwin
goarch: amd64
pkg: juliuszerwick/systems/memory-2/metrics
BenchmarkMetrics/Average_age-4               174           7040887 ns/op
BenchmarkMetrics/Average_payment-4            32          35417759 ns/op
BenchmarkMetrics/Payment_stddev-4             16          69429430 ns/op
PASS
ok      juliuszerwick/systems/memory-2/metrics  11.349s
```


