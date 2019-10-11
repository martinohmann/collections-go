package collections

import (
	"strings"
	"testing"

	"github.com/martinohmann/collections-go/internal/testutil"
)

func BenchmarkImmutableStringFilter1(b *testing.B)   { benchmarkImmutableStringFilter(b, 1) }
func BenchmarkImmutableStringFilter2(b *testing.B)   { benchmarkImmutableStringFilter(b, 2) }
func BenchmarkImmutableStringFilter10(b *testing.B)  { benchmarkImmutableStringFilter(b, 10) }
func BenchmarkImmutableStringFilter100(b *testing.B) { benchmarkImmutableStringFilter(b, 100) }

func benchmarkImmutableStringFilter(b *testing.B, n int) {
	fn := func(item string) bool {
		return !strings.HasPrefix(item, "a")
	}

	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Filter(fn)
	}
}

func BenchmarkImmutableStringPartition1(b *testing.B)   { benchmarkImmutableStringPartition(b, 1) }
func BenchmarkImmutableStringPartition2(b *testing.B)   { benchmarkImmutableStringPartition(b, 2) }
func BenchmarkImmutableStringPartition10(b *testing.B)  { benchmarkImmutableStringPartition(b, 10) }
func BenchmarkImmutableStringPartition100(b *testing.B) { benchmarkImmutableStringPartition(b, 100) }

func benchmarkImmutableStringPartition(b *testing.B, n int) {
	fn := func(item string) bool {
		return strings.HasPrefix(item, "a")
	}

	input := testutil.RandStringSlice(n, 10)

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

func BenchmarkImmutableStringReverse1(b *testing.B)   { benchmarkImmutableStringReverse(b, 1) }
func BenchmarkImmutableStringReverse2(b *testing.B)   { benchmarkImmutableStringReverse(b, 2) }
func BenchmarkImmutableStringReverse10(b *testing.B)  { benchmarkImmutableStringReverse(b, 10) }
func BenchmarkImmutableStringReverse100(b *testing.B) { benchmarkImmutableStringReverse(b, 100) }

func benchmarkImmutableStringReverse(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Reverse()
	}
}

func BenchmarkImmutableStringCut(b *testing.B) {
	input := testutil.RandStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewImmutableString(input).Cut(2, 5)
	}
}

func BenchmarkImmutableStringMap1(b *testing.B)   { benchmarkImmutableStringMap(b, 1) }
func BenchmarkImmutableStringMap2(b *testing.B)   { benchmarkImmutableStringMap(b, 2) }
func BenchmarkImmutableStringMap10(b *testing.B)  { benchmarkImmutableStringMap(b, 10) }
func BenchmarkImmutableStringMap100(b *testing.B) { benchmarkImmutableStringMap(b, 100) }

func benchmarkImmutableStringMap(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

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
