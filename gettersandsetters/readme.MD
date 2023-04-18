# Introduction
Many places in production code require the population of a struct.  
Most population scenarios require some basic logic and the use of other struct members. For example, if member x is true then set member y to something.     
This benchmark test proves which method is the fastest.   

## Benchmarking
* getters to set singular members   
* getters to set multiple members   
* setters to set singular members  
* setters to set multiple members  
* init to return a pre-populated struct  


## Results

• benchmark results were obtained using go version 1.20.3 on Apple M1 (Pro) Silicon  


```
goos: darwin
goarch: arm64
pkg: benchmarks/gettersandsetters
=== RUN   Benchmark_Getters
Benchmark_Getters
Benchmark_Getters-10                    685775350                1.645 ns/op           0 B/op          0 allocs/op
=== RUN   Benchmark_SettersInSequence
Benchmark_SettersInSequence
Benchmark_SettersInSequence-10          693012198                1.734 ns/op           0 B/op          0 allocs/op
=== RUN   Benchmark_SetEverything
Benchmark_SetEverything
Benchmark_SetEverything-10              688615816                1.727 ns/op           0 B/op          0 allocs/op
=== RUN   Benchmark_InitInline
Benchmark_InitInline
Benchmark_InitInline-10                 300817219                4.007 ns/op           0 B/op          0 allocs/op
PASS
ok      benchmarks/gettersandsetters    5.864s


```


1st column shows number of operations made in 100 ms   
2nd column shows nanoseconds per operation   
4th column shows number of heap allocations   

## Conclusion  

Except the (slowest) InitInline function, The speeds are almost equivalent.  
The choice may be down to coding style / unit testing ability / ease of understanding / readability.  
The getters code allows unit testing with basic types that need no knowledge of the struct. With getters, the struct may change name, package, members with maturity of the project but those tests won't need to change with it.   
