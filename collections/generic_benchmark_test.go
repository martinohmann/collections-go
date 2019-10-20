package collections

import (
	"strings"
	"testing"

	"github.com/martinohmann/collections-go/internal/testutil"
)

func BenchmarkGenericNew(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input)
	}
}

func BenchmarkGenericFilter1(b *testing.B)   { benchmarkGenericFilter(b, 1) }
func BenchmarkGenericFilter2(b *testing.B)   { benchmarkGenericFilter(b, 2) }
func BenchmarkGenericFilter10(b *testing.B)  { benchmarkGenericFilter(b, 10) }
func BenchmarkGenericFilter100(b *testing.B) { benchmarkGenericFilter(b, 100) }

func benchmarkGenericFilter(b *testing.B, n int) {
	fn := func(item interface{}) bool {
		return !strings.HasPrefix(item.(string), "a")
	}

	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Filter(fn)
	}
}

func BenchmarkGenericPartition1(b *testing.B)   { benchmarkGenericPartition(b, 1) }
func BenchmarkGenericPartition2(b *testing.B)   { benchmarkGenericPartition(b, 2) }
func BenchmarkGenericPartition10(b *testing.B)  { benchmarkGenericPartition(b, 10) }
func BenchmarkGenericPartition100(b *testing.B) { benchmarkGenericPartition(b, 100) }

func benchmarkGenericPartition(b *testing.B, n int) {
	fn := func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "a")
	}

	input := testutil.RandStringSlice(n, 10)

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

func BenchmarkGenericReverse1(b *testing.B)   { benchmarkGenericReverse(b, 1) }
func BenchmarkGenericReverse2(b *testing.B)   { benchmarkGenericReverse(b, 2) }
func BenchmarkGenericReverse10(b *testing.B)  { benchmarkGenericReverse(b, 10) }
func BenchmarkGenericReverse100(b *testing.B) { benchmarkGenericReverse(b, 100) }

func benchmarkGenericReverse(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Reverse()
	}
}

func BenchmarkGenericCopy(b *testing.B) {
	input := testutil.RandStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Copy()
	}
}

func BenchmarkGenericCut(b *testing.B) {
	input := testutil.RandStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New(input).Cut(2, 5)
	}
}

func BenchmarkGenericMap1(b *testing.B)   { benchmarkGenericMap(b, 1) }
func BenchmarkGenericMap2(b *testing.B)   { benchmarkGenericMap(b, 2) }
func BenchmarkGenericMap10(b *testing.B)  { benchmarkGenericMap(b, 10) }
func BenchmarkGenericMap100(b *testing.B) { benchmarkGenericMap(b, 100) }

func benchmarkGenericMap(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

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
