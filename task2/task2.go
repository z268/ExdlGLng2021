package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const repeats_treshold = 4

func pack(input string) string {
	var (
		prevChar rune
		repeats  int
		result   string
	)

	flushRepeats := func(char rune) {
		if repeats > repeats_treshold {
			result += fmt.Sprintf("#%v#%c", repeats, prevChar)
		} else {
			result += strings.Repeat(string(prevChar), repeats)
		}
		prevChar, repeats = char, 1
	}

	for _, char := range input {
		if char == prevChar {
			repeats += 1
		} else {
			flushRepeats(char)
		}
	}
	flushRepeats(0)

	return result
}

func unpack(input_raw string) string {
	var (
		repeatsStr, result string
		isInsideHashes     bool
	)

	for _, char := range []rune(input_raw) {
		if char == '#' {
			isInsideHashes = !isInsideHashes
			continue
		}

		if (len(repeatsStr) > 0) && !isInsideHashes {
			repeats, _ := strconv.Atoi(repeatsStr)
			result += strings.Repeat(string(char), repeats)
			repeatsStr = ""
		} else if isInsideHashes {
			repeatsStr += string(char)
		} else {
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Print("Enter string: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	if strings.ContainsRune(input, '#') {
		fmt.Printf("Unpacked: %v", unpack(input))
	} else {
		fmt.Printf("Packed: %v", pack(input))
	}
}
