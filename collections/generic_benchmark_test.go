package collections

import (
	"strings"
	"testing"
)

func BenchmarkGenericFilter10_5(b *testing.B) {
	benchmarkGenericFilter(b, 10, 5)
}

func BenchmarkGenericFilter1000_20(b *testing.B) {
	benchmarkGenericFilter(b, 1000, 20)
}

func benchmarkGenericFilter(b *testing.B, n, strlen int) {
	fn := func(item interface{}) bool {
		return !strings.HasPrefix(item.(string), "a")
	}

	input := randomStringSlice(5, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Filter(fn)
	}
}

func BenchmarkGenericPartition10_5(b *testing.B) {
	benchmarkGenericPartition(b, 10, 5)
}

func BenchmarkGenericPartition1000_20(b *testing.B) {
	benchmarkGenericPartition(b, 1000, 20)
}

func benchmarkGenericPartition(b *testing.B, n, strlen int) {
	fn := func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "a")
	}

	input := randomStringSlice(5, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Partition(fn)
	}
}

func BenchmarkGenericInsertItem(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).InsertItem("a", 0)
	}
}

func BenchmarkGenericReverse(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Reverse()
	}
}

func BenchmarkGenericCopy(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Copy()
	}
}

func BenchmarkGenericCut(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Cut(2, 5)
	}
}

func BenchmarkGenericMap(b *testing.B) {
	input := randomStringSlice(100, 10)

	fn := func(item interface{}) interface{} {
		return item
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Map(fn)
	}
}

func BenchmarkGenericPrepend(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Prepend("a")
	}
}
