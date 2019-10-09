package collections

import (
	"strings"
	"testing"
)

func BenchmarkImmutableStringFilter10_5(b *testing.B) {
	benchmarkImmutableStringFilter(b, 10, 5)
}

func BenchmarkImmutableStringFilter1000_20(b *testing.B) {
	benchmarkImmutableStringFilter(b, 1000, 20)
}

func benchmarkImmutableStringFilter(b *testing.B, n, strlen int) {
	fn := func(item string) bool {
		return !strings.HasPrefix(item, "a")
	}

	input := randomStringSlice(5, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Filter(fn)
	}
}

func BenchmarkImmutableStringPartition10_5(b *testing.B) {
	benchmarkImmutableStringPartition(b, 10, 5)
}

func BenchmarkImmutableStringPartition1000_20(b *testing.B) {
	benchmarkImmutableStringPartition(b, 1000, 20)
}

func benchmarkImmutableStringPartition(b *testing.B, n, strlen int) {
	fn := func(item string) bool {
		return strings.HasPrefix(item, "a")
	}

	input := randomStringSlice(5, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Partition(fn)
	}
}

func BenchmarkImmutableStringInsertItem(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).InsertItem("a", 0)
	}
}

func BenchmarkImmutableStringReverse(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Reverse()
	}
}

func BenchmarkImmutableStringCut(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Cut(2, 5)
	}
}

func BenchmarkImmutableStringMap(b *testing.B) {
	input := randomStringSlice(100, 10)

	fn := func(item string) string {
		return item
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Map(fn)
	}
}

func BenchmarkImmutableStringPrepend(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Prepend("a")
	}
}
