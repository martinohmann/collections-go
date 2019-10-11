package collections

import (
	"strings"
	"testing"

	"github.com/martinohmann/collections-go/internal/testutil"
)

func BenchmarkStringFilter1(b *testing.B)   { benchmarkStringFilter(b, 1) }
func BenchmarkStringFilter2(b *testing.B)   { benchmarkStringFilter(b, 2) }
func BenchmarkStringFilter10(b *testing.B)  { benchmarkStringFilter(b, 10) }
func BenchmarkStringFilter100(b *testing.B) { benchmarkStringFilter(b, 100) }

func benchmarkStringFilter(b *testing.B, n int) {
	fn := func(item string) bool {
		return !strings.HasPrefix(item, "a")
	}

	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Filter(fn)
	}
}

func BenchmarkStringPartition1(b *testing.B)   { benchmarkStringPartition(b, 1) }
func BenchmarkStringPartition2(b *testing.B)   { benchmarkStringPartition(b, 2) }
func BenchmarkStringPartition10(b *testing.B)  { benchmarkStringPartition(b, 10) }
func BenchmarkStringPartition100(b *testing.B) { benchmarkStringPartition(b, 100) }

func benchmarkStringPartition(b *testing.B, n int) {
	fn := func(item string) bool {
		return strings.HasPrefix(item, "a")
	}

	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Partition(fn)
	}
}

func BenchmarkStringInsertItem(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).InsertItem("a", 0)
	}
}

func BenchmarkStringReverse1(b *testing.B)   { benchmarkStringReverse(b, 1) }
func BenchmarkStringReverse2(b *testing.B)   { benchmarkStringReverse(b, 2) }
func BenchmarkStringReverse10(b *testing.B)  { benchmarkStringReverse(b, 10) }
func BenchmarkStringReverse100(b *testing.B) { benchmarkStringReverse(b, 100) }

func benchmarkStringReverse(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Reverse()
	}
}

func BenchmarkStringCopy(b *testing.B) {
	input := testutil.RandStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Copy()
	}
}

func BenchmarkStringCut(b *testing.B) {
	input := testutil.RandStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Cut(2, 5)
	}
}

func BenchmarkStringMap1(b *testing.B)   { benchmarkStringMap(b, 1) }
func BenchmarkStringMap2(b *testing.B)   { benchmarkStringMap(b, 2) }
func BenchmarkStringMap10(b *testing.B)  { benchmarkStringMap(b, 10) }
func BenchmarkStringMap100(b *testing.B) { benchmarkStringMap(b, 100) }

func benchmarkStringMap(b *testing.B, n int) {
	input := testutil.RandStringSlice(n, 10)

	fn := func(item string) string {
		return item
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Map(fn)
	}
}

func BenchmarkStringPrepend(b *testing.B) {
	input := []string{"a"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Prepend("a")
	}
}
