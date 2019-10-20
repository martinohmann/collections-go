package collections

import (
	"strings"
	"testing"

	"github.com/martinohmann/collections-go/internal/testutil"
)

func BenchmarkImmutableGenericNew(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input)
	}
}

func BenchmarkImmutableGenericFilter1(b *testing.B)   { benchmarkImmutableGenericFilter(b, 1) }
func BenchmarkImmutableGenericFilter2(b *testing.B)   { benchmarkImmutableGenericFilter(b, 2) }
func BenchmarkImmutableGenericFilter10(b *testing.B)  { benchmarkImmutableGenericFilter(b, 10) }
func BenchmarkImmutableGenericFilter100(b *testing.B) { benchmarkImmutableGenericFilter(b, 100) }

func benchmarkImmutableGenericFilter(b *testing.B, n int) {
	fn := func(item interface{}) bool {
		return !strings.HasPrefix(item.(string), "a")
	}

	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Filter(fn)
	}
}

func BenchmarkImmutableGenericPartition1(b *testing.B)   { benchmarkImmutableGenericPartition(b, 1) }
func BenchmarkImmutableGenericPartition2(b *testing.B)   { benchmarkImmutableGenericPartition(b, 2) }
func BenchmarkImmutableGenericPartition10(b *testing.B)  { benchmarkImmutableGenericPartition(b, 10) }
func BenchmarkImmutableGenericPartition100(b *testing.B) { benchmarkImmutableGenericPartition(b, 100) }

func benchmarkImmutableGenericPartition(b *testing.B, n int) {
	fn := func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "a")
	}

	input := testutil.RandStringSlice(n, 10)

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

func BenchmarkImmutableGenericReverse1(b *testing.B)   { benchmarkImmutableGenericReverse(b, 1) }
func BenchmarkImmutableGenericReverse2(b *testing.B)   { benchmarkImmutableGenericReverse(b, 2) }
func BenchmarkImmutableGenericReverse10(b *testing.B)  { benchmarkImmutableGenericReverse(b, 10) }
func BenchmarkImmutableGenericReverse100(b *testing.B) { benchmarkImmutableGenericReverse(b, 100) }

func benchmarkImmutableGenericReverse(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Reverse()
	}
}

func BenchmarkImmutableGenericCut(b *testing.B) {
	input := testutil.RandStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutable(input).Cut(2, 5)
	}
}

func BenchmarkImmutableGenericMap1(b *testing.B)   { benchmarkImmutableGenericMap(b, 1) }
func BenchmarkImmutableGenericMap2(b *testing.B)   { benchmarkImmutableGenericMap(b, 2) }
func BenchmarkImmutableGenericMap10(b *testing.B)  { benchmarkImmutableGenericMap(b, 10) }
func BenchmarkImmutableGenericMap100(b *testing.B) { benchmarkImmutableGenericMap(b, 100) }

func benchmarkImmutableGenericMap(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

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
