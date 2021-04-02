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
		prev_char rune
		repeats   int
		result    string
	)

	flushRepeats := func(char rune) {
		if repeats > repeats_treshold {
			result += fmt.Sprintf("#%v#%c", repeats, prev_char)
		} else {
			result += strings.Repeat(string(prev_char), repeats)
		}
		prev_char, repeats = char, 1
	}

	for _, char := range input {
		if char == prev_char {
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
		repeats_str, result string
		inside_hashes       bool
	)

	for _, char := range []rune(input_raw) {
		if char == '#' {
			inside_hashes = !inside_hashes
			continue
		}

		if (len(repeats_str) > 0) && !inside_hashes {
			repeats, _ := strconv.Atoi(repeats_str)
			result += strings.Repeat(string(char), repeats)
			repeats_str = ""
		} else if inside_hashes {
			repeats_str += string(char)
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
