package testutil

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var source = rand.NewSource(time.Now().UnixNano())

// RandStringSlice returns a slice which contains n random strings of length l.
// The strings are generated using RandString.
func RandStringSlice(n, l int) []string {
	s := make([]string, n)

	for i := 0; i < n; i++ {
		s = append(s, RandString(l))
	}

	return s
}

// RandString returns a random string of length l which contains only lowercase
// letters, uppercase letters and numbers.
func RandString(l int) string {
	var sb strings.Builder
	sb.Grow(l)

	for i := 0; i < l; i++ {
		b := source.Int63() % int64(len(alphabet))

		sb.WriteByte(alphabet[b])
	}

	return sb.String()
}
