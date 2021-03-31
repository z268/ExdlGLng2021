package main

import (
	"testing"
)

var (
	big_automorphic_number  = 7109376
	automorphic_numbers     = []int{0, 1, 5, 6, 25, 625, big_automorphic_number}
	non_automorphic_numbers = []int{3, 9, 11, 35, 55, 125, 255}
)

func TestIsAutomorphic(t *testing.T) {
	for _, number := range automorphic_numbers {
		if !isAutomorphic(number) {
			t.Fatalf(`%v is automorphic`, number)
		}
	}

	for _, number := range non_automorphic_numbers {
		if isAutomorphic(number) {
			t.Fatalf(`%v is not automorphic`, number)
		}
	}

}

func BenchmarkIsAutomorphic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = isAutomorphic(big_automorphic_number)
	}
}
