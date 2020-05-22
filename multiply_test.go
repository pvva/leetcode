package leetcode

import (
	"math/rand"
	"testing"
)

// this is a proof that we don't need to optimise multiplication as it is very well done by modern compilers+CPUs

var r int32

func BenchmarkMul10Mul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := rand.Int31n(1e9)
		r = t * 10 // one operation, ~6 ticks
	}
}

func BenchmarkMul10Shift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := rand.Int31n(1e9)
		r = t<<3 + t<<1 // *8 + *2, 3 operations, ~6 ticks in total
	}
}

func BenchmarkMul8Mul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := rand.Int31n(1e9)
		r = t * 8 // one operation, 2 ticks (mul + overflow check) because it is optimised as an argument is a power of 2
	}
}

func BenchmarkMul8Shift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := rand.Int31n(1e9)
		r = t << 3 // *8, 2 ticks (shift + overflow check)
	}
}
