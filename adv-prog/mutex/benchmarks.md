# Benchmarks

Benchmarks with the name `BenchmarkMutex` consider the custom implementation of a mutex, while those with the name `BenchmarkRealMutex` consider the sync.Mutex implementation in the Go source code. 

First run:

```
 go test -bench=.
 goos: darwin
 goarch: amd64
 pkg: juliuszerwick/adv-prog/mutex
 BenchmarkMutex10-4                210512              5787 ns/op
 BenchmarkRealMutex10-4            199804              5353 ns/op
 PASS
 ok      juliuszerwick/adv-prog/mutex    2.570s
 ```


 Second run:

 ```
  go test -bench=.
  goos: darwin
  goarch: amd64
  pkg: juliuszerwick/adv-prog/mutex
  BenchmarkMutex10-4                242846              5505 ns/op
  BenchmarkMutex100-4                31729             48346 ns/op
  BenchmarkMutex1000-4                3775            349057 ns/op
  BenchmarkRealMutex10-4            221100              5440 ns/op
  BenchmarkRealMutex100-4            31028             36306 ns/op
  BenchmarkRealMutex1000-4            3844            348691 ns/op
  PASS
  ok      juliuszerwick/adv-prog/mutex    10.149s

 ```
