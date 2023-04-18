# Introduction 
A recent investigation into a possible memory leak in production code led to the use of PPROF - the inbuilt Golang profiling tool.   

Link to Gopher Con 2019 PPROF demo / tutorial: [Dave Cheney GopherCon 2019](https://www.youtube.com/watch?v=nok0aYiGiYA)  

A heap profile identified a couple of functions that consumed large amounts of **heap** memory.  

## Heap Usage & Performance (Speed) 
Upon close inspection the functions could be rewritten to avoid heap memory usage which at the same time resulted in a performance gain in terms of speed.  
This project shows how the golang benchmarking tool can prove the speed gain.

## Benchmarking pre-allocated slices and maps

using plain int types in slices and maps to avoid any skewing of timings (ie. if the values were strings, we would be including timing of the fmt.Sprintf func or some other func to convert from the loop counter i to a string value).  
Slices have the ability to specify length and capacity, maps do not have the concept of capacity, they just have a size hint but will be reallocated space if required.  
The benchmarks have been orchestrated to build maps (and slices) with 1, 16, and 1024 elements to gain an understanding of the difference in performance across different scenarios.  



### Results

• benchmark results were obtained using go version 1.20.3 on Apple M1 (Pro) Silicon  

```
goos: darwin
goarch: arm64
pkg: benchmarks/preallocation/maps
=== RUN   BenchmarkMaps_UnAllocated1
BenchmarkMaps_UnAllocated1
BenchmarkMaps_UnAllocated1-10           152167297                7.816 ns/op           0 B/op          0 allocs/op
=== RUN   BenchmarkMaps_UnAllocated16
BenchmarkMaps_UnAllocated16
BenchmarkMaps_UnAllocated16-10           2153444               555.4 ns/op           918 B/op          3 allocs/op
=== RUN   BenchmarkMaps_UnAllocated1024
BenchmarkMaps_UnAllocated1024
BenchmarkMaps_UnAllocated1024-10           31984             37280 ns/op           82324 B/op         19 allocs/op
=== RUN   BenchmarkMaps_PreAllocated1
BenchmarkMaps_PreAllocated1
BenchmarkMaps_PreAllocated1-10          127918495                9.373 ns/op           0 B/op          0 allocs/op
=== RUN   BenchmarkMaps_PreAllocated16
BenchmarkMaps_PreAllocated16
BenchmarkMaps_PreAllocated16-10          3933122               307.8 ns/op           580 B/op          1 allocs/op
=== RUN   BenchmarkMaps_PreAllocated1024
BenchmarkMaps_PreAllocated1024
BenchmarkMaps_PreAllocated1024-10         126474              9310 ns/op           40984 B/op          2 allocs/op
PASS
ok      benchmarks/preallocation/maps   10.396s
```

```
goos: darwin
goarch: arm64
pkg: benchmarks/preallocation/slices
=== RUN   BenchmarkSlices_UnAllocated1
BenchmarkSlices_UnAllocated1
BenchmarkSlices_UnAllocated1-10                 83303420                14.36 ns/op            8 B/op          1 allocs/op
=== RUN   BenchmarkSlices_UnAllocated16
BenchmarkSlices_UnAllocated16
BenchmarkSlices_UnAllocated16-10                11260142               103.7 ns/op           248 B/op          5 allocs/op
=== RUN   BenchmarkSlices_UnAllocated1024
BenchmarkSlices_UnAllocated1024
BenchmarkSlices_UnAllocated1024-10                456147              2440 ns/op           25208 B/op         12 allocs/op
=== RUN   BenchmarkSlices_CapacityAllocated1
BenchmarkSlices_CapacityAllocated1
BenchmarkSlices_CapacityAllocated1-10           128151685                9.481 ns/op           8 B/op          1 allocs/op
=== RUN   BenchmarkSlices_CapacityAllocated16
BenchmarkSlices_CapacityAllocated16
BenchmarkSlices_CapacityAllocated16-10          43819875                27.74 ns/op          128 B/op          1 allocs/op
=== RUN   BenchmarkSlices_CapacityAllocated1024
BenchmarkSlices_CapacityAllocated1024
BenchmarkSlices_CapacityAllocated1024-10         1000000              1118 ns/op            8192 B/op          1 allocs/op
=== RUN   BenchmarkSlices_LengthAllocated1
BenchmarkSlices_LengthAllocated1
BenchmarkSlices_LengthAllocated1-10             127952793                9.352 ns/op           8 B/op          1 allocs/op
=== RUN   BenchmarkSlices_LengthAllocated16
BenchmarkSlices_LengthAllocated16
BenchmarkSlices_LengthAllocated16-10            44934633                26.31 ns/op          128 B/op          1 allocs/op
=== RUN   BenchmarkSlices_LengthAllocated1024
BenchmarkSlices_LengthAllocated1024
BenchmarkSlices_LengthAllocated1024-10           1000000              1018 ns/op            8192 B/op          1 allocs/op
PASS
ok      benchmarks/preallocation/slices 13.584s
```

1st column shows number of operations made in 1 second  
2nd column shows nanoseconds per operation (where operation is a func call)


