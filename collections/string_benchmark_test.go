package collections

import (
	"math/rand"
	"strings"
	"testing"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func BenchmarkStringFilter10_5(b *testing.B) {
	benchmarkStringFilter(b, 10, 5)
}

func BenchmarkStringFilter1000_20(b *testing.B) {
	benchmarkStringFilter(b, 1000, 20)
}

func benchmarkStringFilter(b *testing.B, n, strlen int) {
	fn := func(item string) bool {
		return !strings.HasPrefix(item, "a")
	}

	input := randomStringSlice(5, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Filter(fn)
	}
}

func BenchmarkStringPartition10_5(b *testing.B) {
	benchmarkStringPartition(b, 10, 5)
}

func BenchmarkStringPartition1000_20(b *testing.B) {
	benchmarkStringPartition(b, 1000, 20)
}

func benchmarkStringPartition(b *testing.B, n, strlen int) {
	fn := func(item string) bool {
		return strings.HasPrefix(item, "a")
	}

	input := randomStringSlice(5, 10)

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

func BenchmarkStringReverse(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Reverse()
	}
}

func BenchmarkStringCopy(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Copy()
	}
}

func BenchmarkStringCut(b *testing.B) {
	input := randomStringSlice(100, 10)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewString(input).Cut(2, 5)
	}
}

func BenchmarkStringMap(b *testing.B) {
	input := randomStringSlice(100, 10)

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

func randomStringSlice(n, strlen int) []string {
	s := make([]string, n)
	rand.Int63n(int64(len(alphabet) - 1))

	for i := 0; i < n; i++ {
		r := make([]byte, strlen)
		for j := 0; j < strlen; j++ {
			r[j] = alphabet[rand.Int63n(int64(len(alphabet)-1))]
		}

		s[i] = string(r)
	}

	return s
}
