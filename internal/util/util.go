package util

import (
	"math/rand/v2"
)

// Number is a constraint that matches all numeric types.
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Abs returns the absolute value of the given number.
func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// Shuffle randomly shuffles the elements of the input slice and returns a new slice with the shuffled elements.
func Shuffle[T any](s []T) []T {
	results := make([]T, len(s))
	copy(results, s)
	rand.Shuffle(len(results), func(i, j int) {
		results[i], results[j] = results[j], results[i]
	})
	return results
}

// RandomPick randomly selects and returns a single element from the input slice.
func RandomPick[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	idx := rand.IntN(len(s))
	return s[idx], true
}
