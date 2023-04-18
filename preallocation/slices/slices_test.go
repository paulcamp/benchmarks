package slices

import "testing"

func slicesUnallocated(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		makeUnAllocatedSlice(i)
	}
}

func slicesWithCapacity(i int, b *testing.B) {

	for n := 0; n < b.N; n++ {
		makeSliceWithPreAllocatedCapacity(i)
	}
}

func slicesWithLength(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		makeSlicewithPreAllocatedLength(i)
	}
}

func BenchmarkSlices_UnAllocated1(b *testing.B)          { slicesUnallocated(1, b) }
func BenchmarkSlices_UnAllocated16(b *testing.B)         { slicesUnallocated(16, b) }
func BenchmarkSlices_UnAllocated1024(b *testing.B)       { slicesUnallocated(1024, b) }
func BenchmarkSlices_CapacityAllocated1(b *testing.B)    { slicesWithCapacity(1, b) }
func BenchmarkSlices_CapacityAllocated16(b *testing.B)   { slicesWithCapacity(16, b) }
func BenchmarkSlices_CapacityAllocated1024(b *testing.B) { slicesWithCapacity(1024, b) }
func BenchmarkSlices_LengthAllocated1(b *testing.B)      { slicesWithLength(1, b) }
func BenchmarkSlices_LengthAllocated16(b *testing.B)     { slicesWithLength(16, b) }
func BenchmarkSlices_LengthAllocated1024(b *testing.B)   { slicesWithLength(1024, b) }
