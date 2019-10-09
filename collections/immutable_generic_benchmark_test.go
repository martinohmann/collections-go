package collections

import (
	"strings"
	"testing"
)

func BenchmarkImmutableGenericFilter10_5(b *testing.B) {
	benchmarkImmutableGenericFilter(b, 10, 5)
}

func BenchmarkImmutableGenericFilter1000_20(b *testing.B) {
	benchmarkImmutableGenericFilter(b, 1000, 20)
}

func benchmarkImmutableGenericFilter(b *testing.B, n, strlen int) {
	fn := func(item interface{}) bool {
		return !strings.HasPrefix(item.(string), "a")
	}

	input := randomStringSlice(5, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Filter(fn)
	}
}

func BenchmarkImmutableGenericPartition10_5(b *testing.B) {
	benchmarkImmutableGenericPartition(b, 10, 5)
}

func BenchmarkImmutableGenericPartition1000_20(b *testing.B) {
	benchmarkImmutableGenericPartition(b, 1000, 20)
}

func benchmarkImmutableGenericPartition(b *testing.B, n, strlen int) {
	fn := func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "a")
	}

	input := randomStringSlice(5, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Partition(fn)
	}
}

func BenchmarkImmutableGenericInsertItem(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).InsertItem("a", 0)
	}
}

func BenchmarkImmutableGenericReverse(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Reverse()
	}
}

func BenchmarkImmutableGenericCut(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Cut(2, 5)
	}
}

func BenchmarkImmutableGenericMap(b *testing.B) {
	input := randomStringSlice(100, 10)

	fn := func(item interface{}) interface{} {
		return item
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Map(fn)
	}
}

func BenchmarkImmutableGenericPrepend(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Prepend("a")
	}
}
