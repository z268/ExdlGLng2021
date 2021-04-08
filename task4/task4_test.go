package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	encrypted   = "max ctr, ibz, yhq, sxukt tgw fr pheoxl jntvd! max ybox uhqbgz pbstkwl cnfi jnbvder. itvd fr uhq pbma ybox whsxg ebjnhk cnzl."
	expected    = "the jay, pig, fox, zebra and my wolves quack! the five boxing wizards jump quickly. pack my box with five dozen liquor jugs."
	expectedKey = 19
)

func getRandomWords(words []string, limit int) []string {
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return words[:limit]
}

func TestCounter(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	encryptedWords := getWords(encrypted)
	knownWords := getWords(expected)

	fiveKnownWords := getRandomWords(knownWords, 5)
	fmt.Printf("Finding key by 5 random words: %v\n", fiveKnownWords)

	key := findKey(encryptedWords, fiveKnownWords)
	if key != expectedKey {
		t.Fatalf("Incorrect key %v, expected %v\n", key, expectedKey)
	} else {
		fmt.Printf("The key - %v.\n", key)
	}

	decrypted := shiftString(encrypted, -key)
	if decrypted != expected {
		t.Fatalf("Incorrect decyption result.\nGot: %v\nExpected: %v\n", decrypted, expected)
	} else {
		fmt.Printf("Decrypted string:\n%v\n", decrypted)
	}
}
