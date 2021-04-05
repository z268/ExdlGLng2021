package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isAutomorphic(number int) bool {
	squaredNumber := number * number
	divider := 10
	for divider < number {
		divider *= 10
	}
	return squaredNumber% divider == number
}

func readNumber() (int, error) {
	fmt.Print("Is a number automorphic? Enter a number: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strconv.Atoi(strings.Trim(input, "\n\r"))
}

func main() {
	number, err := readNumber()
	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	if isAutomorphic(number) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
