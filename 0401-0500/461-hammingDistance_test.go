package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

var seed = time.Now().UTC().UnixNano()
var hd int

// this is slowest variant
// arithmetical operation 2..4 ticks
// 5 logical operations 1 tick each
// so 7..9 ticks per bit (~8)
func hd1(x int, y int) int {
	c := 0

	for x > 0 || y > 0 {
		c += x&1 ^ y&1
		x >>= 1
		y >>= 1
	}

	return c
}

// this is the fastest variant
// arithmetical operation 2..4 ticks
// 2 logical operations 1 tick each
// so 4..6 ticks per bit (~5)
func hd2(x int, y int) int {
	c := 0
	t := x ^ y

	for t > 0 {
		c, t = c+t&1, t>>1
	}

	return c
}

// this is fast but not the fastest variant
// 2 arithmetical operations 2..4 ticks each
// 1 logical operation 1 tick each
// so 5..9 ticks per bit (~7)
func hd3(x int, y int) int {
	c := 0
	t := x ^ y

	for t > 0 {
		t, c = t&(t-1), c+1
	}

	return c
}

func Benchmark_HammingDistance1(b *testing.B) {
	rand.Seed(seed)
	var d int
	for i := 0; i < b.N; i++ {
		d = hd1(rand.Intn(math.MaxInt32), rand.Intn(math.MaxInt32))
	}

	hd = d
}

func Benchmark_HammingDistance2(b *testing.B) {
	rand.Seed(seed)
	var d int
	for i := 0; i < b.N; i++ {
		d = hd2(rand.Intn(math.MaxInt32), rand.Intn(math.MaxInt32))
	}

	hd = d
}

func Benchmark_HammingDistance3(b *testing.B) {
	rand.Seed(seed)
	var d int
	for i := 0; i < b.N; i++ {
		d = hd3(rand.Intn(math.MaxInt32), rand.Intn(math.MaxInt32))
	}

	hd = d
}
