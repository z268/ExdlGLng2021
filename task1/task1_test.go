package main

import (
	"testing"
)

var (
	bigAutomorphicNumber  = 7109376
	automorphicNumbers    = []int{0, 1, 5, 6, 25, 625, bigAutomorphicNumber}
	nonAutomorphicNumbers = []int{3, 9, 11, 35, 55, 125, 255}
)

func TestIsAutomorphic(t *testing.T) {
	for _, number := range automorphicNumbers {
		if !isAutomorphic(number) {
			t.Fatalf(`%v is automorphic`, number)
		}
	}

	for _, number := range nonAutomorphicNumbers {
		if isAutomorphic(number) {
			t.Fatalf(`%v is not automorphic`, number)
		}
	}

}

func BenchmarkIsAutomorphic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = isAutomorphic(bigAutomorphicNumber)
	}
}
