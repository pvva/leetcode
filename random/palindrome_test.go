package main

import (
	"math/rand"
	"testing"
)

func v1(s string) bool {
	l := len(s)
	for i := 0; i < l>>1; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}

func v2(s string) bool {
	l := 0
	h := len(s) - 1
	for l < h {
		if s[l] != s[h] {
			return false
		}
		l++
		h--
	}

	return true
}

var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func BenchmarkV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1(randString(64))
	}
}

func BenchmarkV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v2(randString(64))
	}
}
