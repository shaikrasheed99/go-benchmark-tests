package factorial

import (
	"testing"
)

func BenchmarkFactorialUsingIteration(b *testing.B) {
	number := 5
	for i := 0; i < b.N; i++ {
		factorialUsingIteration(number)
	}
}

func BenchmarkFactorialUsingRecursion(b *testing.B) {
	number := 5
	for i := 0; i < b.N; i++ {
		factorialUsingRecursion(number)
	}
}
