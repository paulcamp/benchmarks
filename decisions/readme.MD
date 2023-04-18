# Introduction
Many places in production code require decisions - usually checking if a value is a member of a set of values.
This benchmark proves which method is the fastest.  

## Benchmarking plain if / then / else and alternatives
* using plain 'if / then / else logic
* maps can make the code look more elegant. Utilising default values (bool defaults to false) means anything not declared in the map will give false by default.
* case statements can combine all the comparisons into a single line that's more readable




## Results

• benchmark results were obtained using go version 1.20.3 on Apple M1 (Pro) Silicon  

```
goos: darwin
goarch: arm64
pkg: benchmarks/decisions
=== RUN   Benchmark_Original
Benchmark_Original
Benchmark_Original-10                   1000000000               0.3176 ns/op          0 B/op          0 allocs/op
=== RUN   Benchmark_UsingMaps
Benchmark_UsingMaps
Benchmark_UsingMaps-10                  194560927                6.196 ns/op           0 B/op          0 allocs/op
=== RUN   Benchmark_UsingSwitchCase
Benchmark_UsingSwitchCase
Benchmark_UsingSwitchCase-10            1000000000               1.037 ns/op           0 B/op          0 allocs/op
PASS
ok      benchmarks/decisions    3.619s
```


1st column shows number of operations made in 1 second  
2nd column shows nanoseconds per operation (where operation is a func call)  
4th column shows number of heap allocations  
**case statements** win in terms of speed. 