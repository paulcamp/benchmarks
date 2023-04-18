package getters_and_setters

import "testing"

var result MutatingState

func Benchmark_Getters(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		result.EffectiveTotalFund = getEffectiveTotalFund(float64(n), float64(n), n%2 == 0)
		result.Fund95, result.Fund95Tax, result.Fund95Taxed = getFund95Fields(result.EffectiveTotalFund)
		result.Fund100, result.Fund100Tax, result.Fund100Taxed = getFund100Fields(result.EffectiveTotalFund)
	}
}

func Benchmark_SettersInSequence(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		setEffectiveTotalFund(&result, float64(n), float64(n), n%2 == 0) //set single field
		setFund95Fields(&result)                                         //set multiple field
		setFund100Fields(&result)                                        //set multiple field
	}
}

func Benchmark_SetEverything(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		setEverything(&result, float64(n), float64(n), n%2 == 0)
	}
}

func Benchmark_InitInline(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		result = initInLine(float64(n), float64(n), n%2 == 0)
	}
}
