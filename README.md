golang_map_bench
================

```
:~/go/src/github.com/xnum/golang_map_bench$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/xnum/golang_map_bench
BenchmarkNativeMap      2000000000               0.49 ns/op
BenchmarkAtomicMap      1000000000               2.16 ns/op
BenchmarkRWLockMap      20000000                73.0 ns/op
BenchmarkSyncMap        10000000               139 ns/op
PASS
ok      github.com/xnum/golang_map_bench        6.473s
```

