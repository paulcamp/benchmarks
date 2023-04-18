package maps

import (
	"testing"
)

func mapsUnallocated(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		makeUnAllocatedMap(i)
	}
}

func mapsPreAllocated(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		makePreAllocatedMap(i)
	}
}

func BenchmarkMaps_UnAllocated1(b *testing.B)     { mapsUnallocated(1, b) }
func BenchmarkMaps_UnAllocated16(b *testing.B)    { mapsUnallocated(16, b) }
func BenchmarkMaps_UnAllocated1024(b *testing.B)  { mapsUnallocated(1024, b) }
func BenchmarkMaps_PreAllocated1(b *testing.B)    { mapsPreAllocated(1, b) }
func BenchmarkMaps_PreAllocated16(b *testing.B)   { mapsPreAllocated(16, b) }
func BenchmarkMaps_PreAllocated1024(b *testing.B) { mapsPreAllocated(1024, b) }
