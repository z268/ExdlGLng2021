package main

import (
	"testing"
)

var data = map[string]string{
	"один, два - это 2, три один два, много слов: один": "один(3) два(2) это(1) 2(1) три(1) много(1) слов(1)",
	"one two three two one": "one(2) two(2) three(1)",
}

func TestCounter(t *testing.T) {
	for source, expected := range(data) {
		result := getWordStats(source)
		if result != expected {
			t.Fatalf("WordStat counting error of `%v`: `%v` != `%v`", source, result, expected)
		}
	}
}
