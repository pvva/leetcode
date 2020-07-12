package main

import (
	"math/rand"
	"testing"
	"time"
)

func rb1(num uint32) uint32 {
	res := uint32(0)

	l := 0
	r := 31

	for l < r {
		res = res | (num&(1<<l))<<(r-l) | (num&(1<<r))>>(r-l)
		l++
		r--
	}

	return res
}

func rb2(num uint32) uint32 {
	res := uint32(0)

	for i := 0; i < 32; i++ {
		res = (res << 1) | (num & 1)
		num >>= 1
	}

	return res
}

func rb3(num uint32) uint32 {
	res := uint32(0)

	for i := 31; i >= 0; i-- {
		res = res | (((num >> (31 - i)) & 1) << i)
	}

	return res
}

var result uint32
var seed = time.Now().UnixNano()

func Benchmark_rb1(b *testing.B) {
	r := uint32(0)

	rand.Seed(seed)
	for i := 0; i < b.N; i++ {
		r = rb1(rand.Uint32())
	}

	result = r
}

func Benchmark_rb2(b *testing.B) {
	r := uint32(0)

	rand.Seed(seed)
	for i := 0; i < b.N; i++ {
		r = rb2(rand.Uint32())
	}

	result = r
}

func Benchmark_rb3(b *testing.B) {
	r := uint32(0)

	rand.Seed(seed)
	for i := 0; i < b.N; i++ {
		r = rb3(rand.Uint32())
	}

	result = r
}
