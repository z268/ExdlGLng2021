package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	firstLetter  = 'a'
	lastLetter   = 'z'
	alphabetSize = lastLetter - firstLetter + 1
)

func getWords(input string) []string {
	input = strings.ToLower(input)
	r := regexp.MustCompile(`[\w]+`)
	return r.FindAllString(input, -1)
}

func normalize(char int32) int32 {
	return (char + alphabetSize) % alphabetSize
}

func shiftString(input string, key int) string {
	inputChars := []rune(strings.ToLower(input))
	outputChars := make([]rune, 0, len(inputChars))
	for _, char := range inputChars {
		if firstLetter <= char && char <= lastLetter {
			char = normalize(char - firstLetter + int32(key)) + firstLetter
		}
		outputChars = append(outputChars, char)
	}
	return string(outputChars)
}

func getVector(word string) (vector string)  {
	chars := []rune(word)
	for _, char := range chars[1:] {
		vector += string(normalize(char - chars[0]))
	}
	return
}

func findKey(encryptedWords, knownWords []string) (key int) {
	wordsByVector := map[string]string{}
	for _, word := range encryptedWords {
		wordsByVector[getVector(word)] = word
	}

	variants := []int{}
	for _, knownWord := range knownWords {
		vector := getVector(knownWord)
		encryptedWord := []rune(wordsByVector[vector])

		for idx, knownChar := range(knownWord) {
			key = int(normalize(encryptedWord[idx] - knownChar))
			variants = append(variants, key)
		}
	}

	// choose median value of key variants
	sort.Ints(variants)
	return variants[len(variants)/2]
}

func main() {
	encrypted := "Max ctr, ibz, yhq, sxukt tgw fr pheoxl jntvd! Max ybox uhqbgz pbstkwl cnfi jnbvder. Itvd fr uhq pbma ybox whsxg ebjnhk cnzl."
	knownWords := "The Pack fox"

	key := findKey(getWords(encrypted), getWords(knownWords))
	fmt.Println("Found key:", key)
	fmt.Println(shiftString(encrypted, -key))
}
