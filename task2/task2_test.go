package main

import (
	"testing"
)

var data = map[string]string{
	"кооооооордината":               "к#7#ордината",
	"cooooooords":                   "c#7#ords",
	"pack        space":             "pack#8# space",
	"oooooon the sidesssssssssssss": "#6#on the side#13#s",
	"leess theeen fiiiive repeats":  "leess theeen fiiiive repeats",
	"fiiiiive":                      "f#5#ive",
	"twentyyyyyyyyyyyyyyyyyyyy":     "twent#20#y",
}

func TestPacking(t *testing.T) {
	for source, expected := range data {
		packed := pack(source)
		if packed != expected {
			t.Fatalf(`Packing %v error: %v != %v`, source, packed, expected)
		}
	}
}

func TestUnpacking(t *testing.T) {
	for expected, source := range data {
		unpacked := unpack(source)
		if unpacked != expected {
			t.Fatalf(`Unpacking %v error: %v != %v`, source, unpacked, expected)
		}
	}
}
