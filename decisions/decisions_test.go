package decisions

import "testing"

func Benchmark_Original(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		numberInside100 := n % 100
		isInKnownSet(numberInside100)
	}
}

func Benchmark_UsingMaps(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		numberInside100 := n % 100
		isInMap(numberInside100)
	}
}

func Benchmark_UsingSwitchCase(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		numberInside100 := n % 100
		isInSwitchCase(numberInside100)
	}
}
